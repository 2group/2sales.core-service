package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/2group/2sales.core-service/internal/grpc"
	employeev1 "github.com/2group/2sales.core-service/pkg/gen/go/employee"
	"github.com/2group/2sales.core-service/pkg/json"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type EmployeeHandler struct {
	employee *grpc.EmployeeClient
}

func NewEmployeeHandler(employee *grpc.EmployeeClient) *EmployeeHandler {
	return &EmployeeHandler{employee: employee}
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "employee_handler").
		Str("method", "CreateEmployee").
		Logger()

	log.Info().Msg("request_received")

	req := &employeev1.CreateEmployeeRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, errors.New("invalid payload"))
		return
	}

	log.Info().Interface("request", req).Msg("calling_employee_microservice")

	resp, err := h.employee.Api.CreateEmployee(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int64("employee_id", resp.GetEmployee().GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *EmployeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "employee_handler").
		Str("method", "GetEmployee").
		Logger()

	log.Info().Msg("request_received")

	employeeIDStr := chi.URLParam(r, "employee_id")
	employeeID, err := strconv.ParseInt(employeeIDStr, 10, 64)
	if err != nil {
		log.Error().Str("employee_id", employeeIDStr).Err(err).Msg("invalid_id")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("employee_id", employeeID).Msg("calling_employee_microservice")

	req := &employeev1.GetEmployeeRequest{
		Lookup: &employeev1.GetEmployeeRequest_Id{Id: employeeID},
	}
	resp, err := h.employee.Api.GetEmployee(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "employee_handler").
		Str("method", "UpdateEmployee").
		Logger()

	log.Info().Msg("request_received")

	employeeIDStr := chi.URLParam(r, "employee_id")
	employeeID, err := strconv.ParseInt(employeeIDStr, 10, 64)
	if err != nil {
		log.Error().Str("employee_id", employeeIDStr).Err(err).Msg("invalid_id")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &employeev1.UpdateEmployeeRequest{
		Employee: &employeev1.Employee{Id: &employeeID},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Interface("request", req).Msg("calling_employee_microservice")

	resp, err := h.employee.Api.UpdateEmployee(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int64("employee_id", employeeID).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *EmployeeHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "employee_handler").
		Str("method", "CreateRole").
		Logger()

	log.Info().Msg("request_received")

	req := &employeev1.CreateRoleRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Interface("request", req).Msg("calling_employee_microservice")

	resp, err := h.employee.Api.CreateRole(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int64("role_id", resp.GetRole().GetId()).Msg("succeeded")
	json.WriteJSON(w, http.StatusCreated, resp)
}

func (h *EmployeeHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "employee_handler").
		Str("method", "UpdateRole").
		Logger()

	log.Info().Msg("request_received")

	roleIDStr := chi.URLParam(r, "role_id")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		log.Error().Str("role_id", roleIDStr).Err(err).Msg("invalid_id")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &employeev1.UpdateRoleRequest{
		Role: &employeev1.Role{Id: &roleID},
	}
	if err := json.ParseJSON(r, req); err != nil {
		log.Error().Err(err).Msg("invalid_payload")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Interface("request", req).Msg("calling_employee_microservice")

	resp, err := h.employee.Api.UpdateRole(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int64("role_id", roleID).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *EmployeeHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "employee_handler").
		Str("method", "DeleteRole").
		Logger()

	log.Info().Msg("request_received")

	roleIDStr := chi.URLParam(r, "role_id")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		log.Error().Str("role_id", roleIDStr).Err(err).Msg("invalid_id")
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	log.Info().Int64("role_id", roleID).Msg("calling_employee_microservice")

	req := &employeev1.DeleteRoleRequest{Id: roleID}
	resp, err := h.employee.Api.DeleteRole(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int64("role_id", roleID).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *EmployeeHandler) ListRoles(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "employee_handler").
		Str("method", "ListRoles").
		Logger()

	log.Info().Msg("request_received")

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

	log.Info().Int("limit", limit).Int("offset", offset).Msg("calling_employee_microservice")

	req := &employeev1.ListRolesRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	resp, err := h.employee.Api.ListRoles(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int("roles_count", len(resp.Roles)).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *EmployeeHandler) ListEmployees(w http.ResponseWriter, r *http.Request) {
	log := zerolog.Ctx(r.Context()).With().
		Str("component", "employee_handler").
		Str("method", "ListEmployees").
		Logger()

	log.Info().Msg("request_received")

	orgIDStr := r.URL.Query().Get("organization_id")
	searchText := r.URL.Query().Get("search_text")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	var orgID int64
	var err error
	if orgIDStr != "" {
		orgID, err = strconv.ParseInt(orgIDStr, 10, 64)
		if err != nil {
			log.Error().Str("organization_id", orgIDStr).Err(err).Msg("invalid_organization_id")
			json.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}

	limit := 5
	if limitStr != "" {
		if limitParsed, err := strconv.Atoi(limitStr); err == nil && limitParsed > 0 {
			limit = limitParsed
		}
	}

	offset := 0
	if offsetStr != "" {
		if offsetParsed, err := strconv.Atoi(offsetStr); err == nil && offsetParsed >= 0 {
			offset = offsetParsed
		}
	}

	log.Info().
		Int64("organization_id", orgID).
		Str("search_text", searchText).
		Int("limit", limit).
		Int("offset", offset).
		Msg("calling_employee_microservice")

	req := &employeev1.ListEmployeesRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	if orgIDStr != "" {
		req.OrganizationId = &orgID
	}

	if searchText != "" {
		req.SearchText = &searchText
	}

	resp, err := h.employee.Api.ListEmployees(r.Context(), req)
	if err != nil {
		log.Error().Err(err).Msg("gRPC_call_failed")
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Info().Int("employees_count", len(resp.Employees)).Msg("succeeded")
	json.WriteJSON(w, http.StatusOK, resp)
}
