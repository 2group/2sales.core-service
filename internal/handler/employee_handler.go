package handler

import (
	"errors"
	"github.com/2group/2sales.core-service/pkg/middleware"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	employeev1 "github.com/2group/2sales.core-service/pkg/gen/go/employee"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
)

type EmployeeHandler struct {
	employee *grpc.EmployeeClient
}

func NewEmployeeHandler(employee *grpc.EmployeeClient) *EmployeeHandler {
	return &EmployeeHandler{
		employee: employee,
	}
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "employee_handler",
		"method", "CreateEmployee",
	)

	log.Info("request_received")

	req := &employeev1.CreateEmployeeRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid payload"))
		return
	}

	log.Info("calling_employee_microservice", "request", req)
	resp, err := h.employee.Api.CreateEmployee(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded", "employee_id", resp.GetEmployee().GetId())
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "employee_handler",
		"method", "GetEmployee",
	)

	log.Info("request_received")

	employeeIDStr := chi.URLParam(r, "employee_id")
	employeeID, err := strconv.ParseInt(employeeIDStr, 10, 64)
	if err != nil {
		log.Error("invalid_id", "employee_id", employeeIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("calling_employee_microservice", "employee_id", employeeID)

	req := &employeev1.GetEmployeeRequest{
		Lookup: &employeev1.GetEmployeeRequest_Id{Id: employeeID},
	}
	resp, err := h.employee.Api.GetEmployee(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "employee_handler",
		"method", "UpdateEmployee",
	)

	log.Info("request_received")

	employeeIDStr := chi.URLParam(r, "employee_id")
	employeeID, err := strconv.ParseInt(employeeIDStr, 10, 64)
	if err != nil {
		log.Error("invalid_id", "employee_id", employeeIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &employeev1.UpdateEmployeeRequest{
		Employee: &employeev1.Employee{Id: &employeeID},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("calling_employee_microservice", "request", req)
	resp, err := h.employee.Api.UpdateEmployee(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded", "employee_id", employeeID)
	json.WriteJSON(w, http.StatusOK, resp)
}

//func (h *EmployeeHandler) ListRole(w http.ResponseWriter, r *http.Request) {
//	h.log.Info("Received request to list Role")
//	req := &employeev1.ListRoleRequest{}
//	if err := json.ParseJSON(r, req); err != nil {
//		h.log.Error("Failed to parse request JSON", "error", err)
//		json.WriteError(w, http.StatusBadRequest, err)
//		return
//	}
//	h.log.Info("Parsed request JSON successfully", "request", req)
//
//	response, err := h.employee.Api.ListRole(r.Context(), req)
//	if err != nil {
//		h.log.Error("Error listing role", "error", err)
//		json.WriteError(w, http.StatusBadRequest, err)
//		return
//	}
//	h.log.Info("Role listed successfully", "response", response)
//
//	json.WriteJSON(w, http.StatusOK, response)
//	h.log.Info("Response sent", "status", http.StatusOK)
//}

func (h *EmployeeHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "employee_handler",
		"method", "CreateRole",
	)
	log.Info("request_received")

	req := &employeev1.CreateRoleRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("calling_employee_microservice", "request", req)
	resp, err := h.employee.Api.CreateRole(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded", "role_id", resp.GetRole().GetId())
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *EmployeeHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "employee_handler",
		"method", "UpdateRole",
	)

	log.Info("request_received")

	roleIDStr := chi.URLParam(r, "role_id")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		log.Error("invalid_id", "role_id", roleIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &employeev1.UpdateRoleRequest{
		Role: &employeev1.Role{Id: &roleID},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info("calling_employee_microservice", "request", req)
	resp, err := h.employee.Api.UpdateRole(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded", "role_id", roleID)
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *EmployeeHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "employee_handler",
		"method", "DeleteRole",
	)

	log.Info("request_received")

	roleIDStr := chi.URLParam(r, "role_id")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		log.Error("invalid_id", "role_id", roleIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &employeev1.DeleteRoleRequest{Id: roleID}
	log.Info("calling_employee_microservice", "role_id", roleID)

	resp, err := h.employee.Api.DeleteRole(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded", "role_id", roleID)
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *EmployeeHandler) ListRoles(w http.ResponseWriter, r *http.Request) {
	log := middleware.LoggerFromContext(r.Context()).With(
		"component", "employee_handler",
		"method", "ListRoles",
	)

	log.Info("request_received")

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 20
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	req := &employeev1.ListRolesRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	log.Info("calling_employee_microservice", "limit", limit, "offset", offset)
	resp, err := h.employee.Api.ListRoles(r.Context(), req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info("succeeded", "roles_count", len(resp.Roles))
	json.WriteJSON(w, http.StatusOK, resp)
}
