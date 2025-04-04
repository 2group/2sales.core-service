package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/2group/2sales.core-service/internal/grpc"
	customerv1 "github.com/2group/2sales.core-service/pkg/gen/go/customer"
	"github.com/2group/2sales.core-service/pkg/json"
)

type CustomerHandler struct {
	log      *slog.Logger
	customer *grpc.CustomerClient
}

func NewCustomerHandler(log *slog.Logger, customer *grpc.CustomerClient) *CustomerHandler {
	return &CustomerHandler{
		log:      log,
		customer: customer,
	}
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request to create customer")

	req := &customerv1.CreateCustomerRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		fmt.Println("Failed to parse request JSON:", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println("Parsed request JSON successfully:", req)

	response, err := h.customer.Api.CreateCustomer(r.Context(), req)
	if err != nil {
		fmt.Println("Error creating customer:", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println("Customer created successfully:", response)

	json.WriteJSON(w, http.StatusCreated, response)
	fmt.Println("Response sent with status 201")
}
