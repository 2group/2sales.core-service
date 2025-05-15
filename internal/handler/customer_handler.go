package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/2group/2sales.core-service/internal/grpc"
	customerv1 "github.com/2group/2sales.core-service/pkg/gen/go/customer"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/2group/2sales.core-service/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type CustomerHandler struct {
	customer *grpc.CustomerClient
}

func NewCustomerHandler(customer *grpc.CustomerClient) *CustomerHandler {
	return &CustomerHandler{customer: customer}
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "CreateCustomer").
		Logger()

	log.Info().Msg("request_received")

	req := &customerv1.CreateCustomerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug().Interface("request", req).Msg("calling_customer_service")

	resp, err := h.customer.Api.CreateCustomer(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("customer_id", resp.Customer.GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "GetCustomer").
		Logger()

	log.Info().Msg("request_received")

	customerIDStr := chi.URLParam(r, "customer_id")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		log.Error().Str("customer_id", customerIDStr).Err(err).Msg("invalid_customer_id")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid customer_id"))
		return
	}

	req := &customerv1.GetCustomerRequest{
		Lookup: &customerv1.GetCustomerRequest_Id{Id: customerID},
	}

	var paths []string
	if strings.EqualFold(r.URL.Query().Get("include_loyalty_level"), "true") {
		paths = append(paths, "loyalty_level")
	}
	if strings.EqualFold(r.URL.Query().Get("include_email"), "true") {
		paths = append(paths, "email")
	}
	if strings.EqualFold(r.URL.Query().Get("include_phone_number"), "true") {
		paths = append(paths, "phone_number")
	}

	if len(paths) > 0 {
		req.FieldMask = &fieldmaskpb.FieldMask{Paths: paths}
		log.Debug().Strs("field_mask.paths", paths).Msg("using field mask")
	}

	log.Debug().Int64("customer_id", customerID).Msg("calling_customer_service")

	resp, err := h.customer.Api.GetCustomer(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusNotFound, err)
		return
	}

	log.Info().Int64("customer_id", resp.Customer.GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *CustomerHandler) GetMyCustomer(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "GetMyCustomer").
		Logger()

	log.Info().Msg("request_received")

	customerID, ok := middleware.GetCustomerID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &customerv1.GetCustomerRequest{
		Lookup: &customerv1.GetCustomerRequest_Id{Id: customerID},
	}

	var paths []string
	if strings.EqualFold(r.URL.Query().Get("include_loyalty_level"), "true") {
		paths = append(paths, "loyalty_level")
	}
	if strings.EqualFold(r.URL.Query().Get("include_email"), "true") {
		paths = append(paths, "email")
	}
	if strings.EqualFold(r.URL.Query().Get("include_phone_number"), "true") {
		paths = append(paths, "phone_number")
	}

	if len(paths) > 0 {
		req.FieldMask = &fieldmaskpb.FieldMask{Paths: paths}
		log.Debug().Strs("field_mask.paths", paths).Msg("using field mask")
	}

	log.Debug().Int64("customer_id", customerID).Msg("calling_customer_service")

	resp, err := h.customer.Api.GetCustomer(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusNotFound, err)
		return
	}

	log.Info().Int64("customer_id", resp.Customer.GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "DeleteCustomer").
		Logger()

	log.Info().Msg("request_received")

	customerIDStr := chi.URLParam(r, "customer_id")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		log.Error().Str("customer_id", customerIDStr).Err(err).Msg("invalid_customer_id")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid customer_id"))
		return
	}

	req := &customerv1.DeleteCustomerRequest{Id: customerID}
	log.Debug().Int64("customer_id", customerID).Msg("calling_customer_service")

	resp, err := h.customer.Api.DeleteCustomer(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("customer_id", customerID).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *CustomerHandler) PartialUpdateCustomer(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "PartialUpdateCustomer").
		Logger()

	log.Info().Msg("request_received")

	customerIDStr := chi.URLParam(r, "customer_id")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		log.Error().Str("customer_id", customerIDStr).Err(err).Msg("invalid_customer_id")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid customer_id"))
		return
	}

	req := &customerv1.PartialUpdateCustomerRequest{
		Customer: &customerv1.Customer{Id: &customerID},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug().Int64("customer_id", customerID).Msg("calling_customer_service")

	resp, err := h.customer.Api.PartialUpdateCustomer(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("customer_id", resp.Customer.GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "UpdateCustomer").
		Logger()

	log.Info().Msg("request_received")

	customerIDStr := chi.URLParam(r, "customer_id")
	customerID, err := strconv.ParseInt(customerIDStr, 10, 64)
	if err != nil {
		log.Error().Str("customer_id", customerIDStr).Err(err).Msg("invalid_customer_id")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid customer_id"))
		return
	}

	req := &customerv1.UpdateCustomerRequest{
		Customer: &customerv1.Customer{Id: &customerID},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if req.Customer.UserId == nil {
		log.Error().Err(err).Msg("user_id must not be empty")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	log.Debug().Int64("customer_id", customerID).Msg("calling_customer_service")
	resp, err := h.customer.Api.UpdateCustomer(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("customer_id", resp.Customer.GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *CustomerHandler) UpdateMyCustomer(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "UpdateMyCustomer").
		Logger()

	log.Info().Msg("request_received")

	customerId, ok := middleware.GetCustomerID(r)
	if !ok {
		json.WriteError(w, http.StatusBadRequest, fmt.Errorf("Unauthorized"))
		return
	}

	req := &customerv1.UpdateCustomerRequest{
		Customer: &customerv1.Customer{Id: &customerId},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Debug().Int64("customer_id", customerId).Msg("calling_customer_service")

	resp, err := h.customer.Api.UpdateCustomer(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("customer_id", resp.Customer.GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *CustomerHandler) ListCustomers(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "customer_handler").
		Str("method", "ListCustomers").
		Logger()

	log.Info().Msg("request_received")

	q := r.URL.Query()

	var (
		limit             int32 = 20
		offset            int32 = 0
		organizationID    *int64
		loyaltyLevelID    *int64
		searchText        *string
		phoneNumberPrefix *string
		createdAtFrom     *string
		createdAtTo       *string
		dateOfBirthFrom   *string
		dateOfBirthTo     *string
	)

	if l := q.Get("limit"); l != "" {
		if val, err := strconv.ParseInt(l, 10, 32); err == nil {
			limit = int32(val)
		}
	}
	if o := q.Get("offset"); o != "" {
		if val, err := strconv.ParseInt(o, 10, 32); err == nil {
			offset = int32(val)
		}
	}
	if org := q.Get("organization_id"); org != "" {
		if val, err := strconv.ParseInt(org, 10, 64); err == nil {
			organizationID = &val
		}
	}
	if loyalty := q.Get("loyalty_level_id"); loyalty != "" {
		if val, err := strconv.ParseInt(loyalty, 10, 64); err == nil {
			loyaltyLevelID = &val
		}
	}
	if s := q.Get("search_text"); s != "" {
		s = strings.TrimSpace(s)
		searchText = &s
	}
	if p := q.Get("phone_number_prefix"); p != "" {
		p = strings.TrimSpace(p)
		phoneNumberPrefix = &p
	}
	if v := q.Get("created_at_from"); v != "" {
		createdAtFrom = &v
	}
	if v := q.Get("created_at_to"); v != "" {
		createdAtTo = &v
	}
	if v := q.Get("date_of_birth_from"); v != "" {
		dateOfBirthFrom = &v
	}
	if v := q.Get("date_of_birth_to"); v != "" {
		dateOfBirthTo = &v
	}

	var fieldMask *fieldmaskpb.FieldMask
	if q.Get("include_loyalty") == "true" {
		fieldMask = &fieldmaskpb.FieldMask{
			Paths: []string{"include_loyalty"},
		}
	}

	req := &customerv1.ListCustomersRequest{
		Limit:             limit,
		Offset:            offset,
		OrganizationId:    organizationID,
		LoyaltyLevelId:    loyaltyLevelID,
		SearchText:        searchText,
		PhoneNumberPrefix: phoneNumberPrefix,
		CreatedAtFrom:     createdAtFrom,
		CreatedAtTo:       createdAtTo,
		DateOfBirthFrom:   dateOfBirthFrom,
		DateOfBirthTo:     dateOfBirthTo,
		FieldMask:         fieldMask,
	}

	log.Debug().Interface("request", req).Msg("calling_customer_service")

	resp, err := h.customer.Api.ListCustomers(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int("customers_count", len(resp.Customers)).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}
