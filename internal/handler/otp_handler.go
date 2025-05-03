package handler

import (
	"errors"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"log/slog"
	"net/http"
	"strings"

	"github.com/2group/2sales.core-service/internal/grpc"
	"github.com/rs/zerolog"

	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"github.com/2group/2sales.core-service/pkg/json"
)

type OtpHandler struct {
	log        *slog.Logger
	otpService *grpc.OtpClient
}

func NewOtpHandler(client *grpc.OtpClient) *OtpHandler {
	return &OtpHandler{
		otpService: client,
	}
}

func (h *OtpHandler) RequestSmsOtp(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "organization_handler").
		Str("method", "GetOrganization").
		Logger()

	log.Info().Msg("request_received")

	var req struct {
		PhoneNumber string `json:"phone_number"`
	}

	if err := json.ParseJSON(r, &req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.otpService.Api.RequestSmsOtp(r.Context(), &userv1.RequestSmsOtpRequest{
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Str("phone_number", req.PhoneNumber).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, map[string]any{
		"ok":      true,
		"message": "Код отправлен",
	})
}

func (h *OtpHandler) VerifySmsOtp(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "otp_handler").
		Str("method", "VerifySmsOtp").
		Logger()

	log.Info().Msg("request_received")

	var reqBody struct {
		PhoneNumber string `json:"phone_number"`
		Code        string `json:"code"`
	}

	if err := json.ParseJSON(r, &reqBody); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// parse optional field mask flags from query
	var paths []string
	if strings.EqualFold(r.URL.Query().Get("include_loyalty"), "true") {
		paths = append(paths, "loyalty_level")
	}
	if strings.EqualFold(r.URL.Query().Get("include_cashback"), "true") {
		paths = append(paths, "cashback_balance")
	}

	req := &userv1.VerifySmsOtpRequest{
		PhoneNumber: reqBody.PhoneNumber,
		Code:        reqBody.Code,
	}

	if len(paths) > 0 {
		req.FieldMask = &fieldmaskpb.FieldMask{Paths: paths}
		log.Debug().Strs("field_mask.paths", paths).Msg("using field mask")
	}

	resp, err := h.otpService.Api.VerifySmsOtp(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusUnauthorized, errors.New("verification failed"))
		return
	}

	log.Info().Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OtpHandler) RequestMailOtp(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "organization_handler").
		Str("method", "GetOrganization").
		Logger()

	log.Info().Msg("request_received")

	var req userv1.RequestMailOtpRequest
	if err := json.ParseProtoJSON(r.Body, &req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.otpService.Api.RequestMailOtp(r.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Str("email", req.Email).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OtpHandler) VerifyMailOtp(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "organization_handler").
		Str("method", "GetOrganization").
		Logger()

	log.Info().Msg("request_received")

	var req userv1.VerifyMailOtpRequest
	if err := json.ParseProtoJSON(r.Body, &req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.otpService.Api.VerifyMailOtp(r.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusUnauthorized, errors.New("verification failed"))
		return
	}

	log.Info().Int64("user_id", *resp.User.Id).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}
