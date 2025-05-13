package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog"
)

type MessageHandler interface {
	HandleMessage(ctx context.Context, msg []byte) error
}

type KafkaConsumer struct {
	logger  zerolog.Logger
	brokers []string
	topic   string
	groupID string
	handler MessageHandler
}

func NewKafkaConsumer(brokers []string, topic, groupID string, handler MessageHandler, logger zerolog.Logger) *KafkaConsumer {
	return &KafkaConsumer{
		logger:  logger.With().Str("component", "kafka_consumer").Logger(),
		brokers: brokers,
		topic:   topic,
		groupID: groupID,
		handler: handler,
	}
}

func (kc *KafkaConsumer) Start(ctx context.Context) error {
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	group, err := sarama.NewConsumerGroup(kc.brokers, kc.groupID, config)
	if err != nil {
		kc.logger.Error().Err(err).Msg("failed_to_create_consumer_group")
		return err
	}

	handler := &consumerGroupHandler{
		logger:  kc.logger,
		handler: kc.handler,
	}

	go func() {
		for {
			if err := group.Consume(ctx, []string{kc.topic}, handler); err != nil {
				kc.logger.Error().Err(err).Msg("consume_error")
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	kc.logger.Info().Str("topic", kc.topic).Msg("kafka_consumer_started")
	return nil
}

type consumerGroupHandler struct {
	handler MessageHandler
	logger  zerolog.Logger
}

func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (h *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		h.logger.Debug().
			Str("topic", msg.Topic).
			Int32("partition", msg.Partition).
			Int64("offset", msg.Offset).
			Msg("message_received")

		if err := h.handler.HandleMessage(context.Background(), msg.Value); err != nil {
			h.logger.Error().Err(err).Msg("handle_message_failed")
			continue
		}

		sess.MarkMessage(msg, "")
	}
	return nil
}
