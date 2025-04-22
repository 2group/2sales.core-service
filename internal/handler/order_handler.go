package handler

import (
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	"github.com/2group/2sales.core-service/pkg/middleware"
)

// OrderHandler handles order-related requests.
type OrderHandler struct {
	log   *slog.Logger
	order *grpc.OrderClient
}

// NewOrderHandler creates a new OrderHandler and enriches the logger with the "component" field.
func NewOrderHandler(log *slog.Logger, order *grpc.OrderClient) *OrderHandler {
	return &OrderHandler{
		log:   log.With("component", "order_handler"),
		order: order,
	}
}

// loggerFor returns a sub-logger enriched with the method name and correlation ID.
func (h *OrderHandler) loggerFor(r *http.Request, method string) *slog.Logger {
	cid, _ := middleware.GetCorrelationID(r.Context())
	return h.log.With("method", method, "correlation_id", cid)
}

// // CreateSubOrder creates a new sub-order.
// func (h *OrderHandler) CreateSubOrder(w http.ResponseWriter, r *http.Request) {
// 	logger := h.loggerFor(r, "CreateSubOrder")
// 	logger.Info("Starting CreateSubOrder")

// 	organizationID, ok := middleware.GetOrganizationID(r)
// 	if !ok {
// 		err := fmt.Errorf("unauthorized: missing organization id")
// 		logger.Error("authorization failed", "error", err)
// 		json.WriteError(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	organizationType, ok := middleware.GetOrganizationType(r)
// 	if !ok {
// 		err := fmt.Errorf("unauthorized: missing organization type")
// 		logger.Error("authorization failed", "error", err)
// 		json.WriteError(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	req := &orderv1.CreateSubOrderRequest{
// 		SubOrder: &orderv1.SubOrder{
// 			FromOrganization: &organizationv1.Organization{},
// 			ToOrganization:   &organizationv1.Organization{},
// 		},
// 	}

// 	if err := json.ParseJSON(r, req); err != nil {
// 		logger.Error("failed to parse request", "error", err)
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	// Set organization IDs based on the organization type.
// 	if organizationType == "retailer" {
// 		req.SubOrder.FromOrganization = &organizationv1.Organization{Id: &organizationID}
// 	} else {
// 		req.SubOrder.FromOrganization.Id = req.SubOrder.ToOrganization.Id
// 		req.SubOrder.ToOrganization = &organizationv1.Organization{Id: &organizationID}
// 	}

// 	logger.Info("calling gRPC CreateSubOrder", "request", req)
// 	response, err := h.order.Api.CreateSubOrder(r.Context(), req)
// 	if err != nil {
// 		logger.Error("gRPC CreateSubOrder call failed", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	logger.Info("CreateSubOrder succeeded", "response", response)
// 	json.WriteJSON(w, http.StatusOK, response)
// }

// // UpdateSubOrder updates an existing sub-order.
// func (h *OrderHandler) UpdateSubOrder(w http.ResponseWriter, r *http.Request) {
// 	logger := h.loggerFor(r, "UpdateSubOrder")
// 	logger.Info("Starting UpdateSubOrder")

// 	suborderIDStr := chi.URLParam(r, "suborder_id")
// 	suborderID, err := strconv.Atoi(suborderIDStr)
// 	if err != nil {
// 		logger.Error("invalid suborder_id", "error", err)
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	req := &orderv1.UpdateSubOrderRequest{Id: int64(suborderID)}
// 	if err := json.ParseJSON(r, &req); err != nil {
// 		logger.Error("failed to parse update request", "error", err)
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	logger.Info("calling gRPC UpdateSubOrder", "request", req)
// 	response, err := h.order.Api.UpdateSubOrder(r.Context(), req)
// 	if err != nil {
// 		logger.Error("gRPC UpdateSubOrder call failed", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	logger.Info("UpdateSubOrder succeeded", "response", response)
// 	json.WriteJSON(w, http.StatusOK, response)
// }

// // GetSubOrder retrieves a sub-order.
// func (h *OrderHandler) GetSubOrder(w http.ResponseWriter, r *http.Request) {
// 	logger := h.loggerFor(r, "GetSubOrder")
// 	logger.Info("Starting GetSubOrder")

// 	suborderIDStr := chi.URLParam(r, "suborder_id")
// 	suborderID, err := strconv.Atoi(suborderIDStr)
// 	if err != nil {
// 		logger.Error("invalid suborder_id", "error", err)
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	req := &orderv1.GetSubOrderRequest{Id: int64(suborderID)}
// 	logger.Info("calling gRPC GetSubOrder", "request", req)
// 	response, err := h.order.Api.GetSubOrder(r.Context(), req)
// 	if err != nil {
// 		logger.Error("gRPC GetSubOrder call failed", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	logger.Info("GetSubOrder succeeded", "response", response)
// 	json.WriteJSON(w, http.StatusOK, response)
// }

// func (h *OrderHandler) ListSubOrder(w http.ResponseWriter, r *http.Request) {
// 	logger := h.loggerFor(r, "ListSubOrder")
// 	logger.Info("Starting ListSubOrder")

// 	organizationID, ok := middleware.GetOrganizationID(r)
// 	if !ok {
// 		err := fmt.Errorf("unauthorized: missing organization id")
// 		logger.Error("authorization failed", "error", err)
// 		json.WriteError(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	query := r.URL.Query()
// 	limit, offset := 10, 0
// 	if limitStr := query.Get("limit"); limitStr != "" {
// 		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
// 			limit = parsedLimit
// 		}
// 	}
// 	if offsetStr := query.Get("offset"); offsetStr != "" {
// 		if parsedOffset, err := strconv.Atoi(offsetStr); err == nil && parsedOffset >= 0 {
// 			offset = parsedOffset
// 		}
// 	}

// 	status := query.Get("status")

// 	req := &orderv1.ListSubOrderRequest{
// 		OrganizationId: organizationID,
// 		Limit:          int64(limit),
// 		Offset:         int64(offset),
// 		Status:         status,
// 	}

// 	logger.Info("calling gRPC ListSubOrder", "request", req)
// 	response, err := h.order.Api.ListSubOrder(r.Context(), req)
// 	if err != nil {
// 		logger.Error("gRPC ListSubOrder call failed", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	logger.Info("ListSubOrder succeeded", "response", response)
// 	json.WriteJSON(w, http.StatusOK, response)
// }

// // GetCart retrieves the shopping cart.
// func (h *OrderHandler) GetCart(w http.ResponseWriter, r *http.Request) {
// 	logger := h.loggerFor(r, "GetCart")
// 	logger.Info("Starting GetCart")

// 	organizationID, ok := middleware.GetOrganizationID(r)
// 	if !ok {
// 		err := fmt.Errorf("unauthorized: missing organization id")
// 		logger.Error("authorization failed", "error", err)
// 		json.WriteError(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	req := &orderv1.ListCartRequest{OrganizationId: organizationID}
// 	logger.Info("calling gRPC GetCart", "request", req)
// 	response, err := h.order.Api.ListCart(r.Context(), req)
// 	if err != nil {
// 		logger.Error("gRPC GetCart call failed", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	logger.Info("GetCart succeeded", "response", response)
// 	json.WriteJSON(w, http.StatusOK, response)
// }

// // AddProductToCart adds a product to the shopping cart.
// func (h *OrderHandler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
// 	logger := h.loggerFor(r, "AddProductToCart")
// 	logger.Info("Starting AddProductToCart")

// 	organizationID, ok := middleware.GetOrganizationID(r)
// 	if !ok {
// 		err := fmt.Errorf("unauthorized: missing organization id")
// 		logger.Error("authorization failed", "error", err)
// 		json.WriteError(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	req := &orderv1.AddProductToCartRequest{
// 		Cart: &orderv1.Cart{
// 			Organization: &organizationv1.Organization{Id: &organizationID},
// 		},
// 	}
// 	if err := json.ParseJSON(r, &req); err != nil {
// 		logger.Error("failed to parse request", "error", err)
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	logger.Info("calling gRPC AddProductToCart", "request", req)
// 	response, err := h.order.Api.AddProductToCart(r.Context(), req)
// 	if err != nil {
// 		logger.Error("gRPC AddProductToCart call failed", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	logger.Info("AddProductToCart succeeded", "response", response)
// 	json.WriteJSON(w, http.StatusOK, response)
// }

// // DeleteProductFromCart deletes a product from the shopping cart.
// func (h *OrderHandler) DeleteProductFromCart(w http.ResponseWriter, r *http.Request) {
// 	logger := h.loggerFor(r, "DeleteProductFromCart")
// 	logger.Info("Starting DeleteProductFromCart")

// 	organizationID, ok := middleware.GetOrganizationID(r)
// 	if !ok {
// 		err := fmt.Errorf("unauthorized: missing organization id")
// 		logger.Error("authorization failed", "error", err)
// 		json.WriteError(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	req := &orderv1.DeleteProductFromCartRequest{
// 		Cart: &orderv1.Cart{
// 			Organization: &organizationv1.Organization{Id: &organizationID},
// 		},
// 	}
// 	if err := json.ParseJSON(r, &req); err != nil {
// 		logger.Error("failed to parse request", "error", err)
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	logger.Info("calling gRPC DeleteProductFromCart", "request", req)
// 	response, err := h.order.Api.DeleteProductFromCart(r.Context(), req)
// 	if err != nil {
// 		logger.Error("gRPC DeleteProductFromCart call failed", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	logger.Info("DeleteProductFromCart succeeded", "response", response)
// 	json.WriteJSON(w, http.StatusOK, response)
// }

// // CreateOrder creates an order.
// func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
// 	logger := h.loggerFor(r, "CreateOrder")
// 	logger.Info("Starting CreateOrder")

// 	organizationID, ok := middleware.GetOrganizationID(r)
// 	if !ok {
// 		err := fmt.Errorf("unauthorized: missing organization id")
// 		logger.Error("authorization failed", "error", err)
// 		json.WriteError(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	req := &orderv1.CreateOrderRequest{OrganizationId: organizationID}
// 	logger.Info("calling gRPC CreateOrder", "request", req)
// 	response, err := h.order.Api.CreateOrder(r.Context(), req)
// 	if err != nil {
// 		logger.Error("gRPC CreateOrder call failed", "error", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	logger.Info("CreateOrder succeeded", "response", response)
// 	json.WriteJSON(w, http.StatusCreated, response)
// }
