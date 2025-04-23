package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	employeev1 "github.com/2group/2sales.core-service/pkg/gen/go/employee"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
)

type EmployeeHandler struct {
	log      *slog.Logger
	employee *grpc.EmployeeClient
}

func NewEmployeeHandler(log *slog.Logger, employee *grpc.EmployeeClient) *EmployeeHandler {
	return &EmployeeHandler{
		log:      log,
		employee: employee,
	}
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to create Employee")

	req := &employeev1.CreateEmployeeRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.employee.Api.CreateEmployee(r.Context(), req)
	if err != nil {
		h.log.Error("Error creating employee", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("employee created successfully", "response", response)

	json.WriteJSON(w, http.StatusCreated, response)
	h.log.Info("Response sent", "status", http.StatusCreated)
}
func (h *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to get Employee")

	employeeIDStr := chi.URLParam(r, "employee_id")

	employeeID, err := strconv.ParseInt(employeeIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid employee_id format", "employee_id", employeeIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &employeev1.GetEmployeeRequest{
		Lookup: &employeev1.GetEmployeeRequest_Id{
			Id: employeeID,
		},
	}

	response, err := h.employee.Api.GetEmployee(r.Context(), req)
	if err != nil {
		h.log.Error("Error getting employee", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("employee retrieved successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}
func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to update Employee")

	employeeIDStr := chi.URLParam(r, "employee_id")

	employeeID, err := strconv.ParseInt(employeeIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid employee_id format", "employee_id", employeeIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &employeev1.UpdateEmployeeRequest{
		Employee: &employeev1.Employee{
			Id: &employeeID,
		},
	}

	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.employee.Api.UpdateEmployee(r.Context(), req)
	if err != nil {
		h.log.Error("Error updating Employee", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("employee updated successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *EmployeeHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to create Role")

	req := &employeev1.CreateRoleRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.employee.Api.CreateRole(r.Context(), req)
	if err != nil {
		h.log.Error("Error creating role", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("role created successfully", "response", response)

	json.WriteJSON(w, http.StatusCreated, response)
	h.log.Info("Response sent", "status", http.StatusCreated)
}

func (h *EmployeeHandler) ListRole(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to list Role")
	req := &employeev1.ListRoleRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.employee.Api.ListRole(r.Context(), req)
	if err != nil {
		h.log.Error("Error listing role", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("Role listed successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *EmployeeHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to update Role")

	roleIDStr := chi.URLParam(r, "role_id")

	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid role_id format", "role_id", roleIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &employeev1.UpdateRoleRequest{
		Role: &employeev1.Role{
			Id: &roleID,
		},
	}

	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.employee.Api.UpdateRole(r.Context(), req)
	if err != nil {
		h.log.Error("Error updating Role", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("role updated successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}
func (h *EmployeeHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to delete Role")

	roleIDStr := chi.URLParam(r, "role_id")

	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		h.log.Error("invalid role_id format", "role_id", roleIDStr, "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &employeev1.DeleteRoleRequest{
		Id: roleID,
	}

	if err := json.ParseJSON(r, req); err != nil {
		h.log.Error("Failed to parse request JSON", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	h.log.Info("Parsed request JSON successfully", "request", req)

	response, err := h.employee.Api.DeleteRole(r.Context(), req)
	if err != nil {
		h.log.Error("Error deleting Role", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
	h.log.Info("role deleted successfully", "response", response)

	json.WriteJSON(w, http.StatusOK, response)
	h.log.Info("Response sent", "status", http.StatusOK)
}

func (h *EmployeeHandler) ListRoles(w http.ResponseWriter, r *http.Request) {
	h.log.Info("Received request to list roles")

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

	resp, err := h.employee.Api.ListRoles(r.Context(), req)
	if err != nil {
		h.log.Error("Error listing roles", "error", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, resp)
}
