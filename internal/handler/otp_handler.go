package handler

import (
	"errors"
	"log/slog"
	"net/http"

	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
	"github.com/2group/2sales.core-service/pkg/json"
)

type OtpHandler struct {
	log        *slog.Logger
	otpService userv1.OtpServiceClient
}

func NewOtpHandler(log *slog.Logger, otpClient userv1.OtpServiceClient) *OtpHandler {
	return &OtpHandler{
		log:        log,
		otpService: otpClient,
	}
}

func (h *OtpHandler) RequestSmsOtp(w http.ResponseWriter, r *http.Request) {
	log := h.log.With("method", "RequestSmsOtp")
	log.Info("request_received")

	var req struct {
		PhoneNumber string `json:"phone_number"`
	}

	if err := json.ParseJSON(r, &req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.otpService.RequestSmsOtp(r.Context(), &userv1.RequestSmsOtpRequest{
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, errors.New("failed to request OTP"))
		return
	}

	log.Info("succeeded", "phone_number", req.PhoneNumber)
	json.WriteJSON(w, http.StatusOK, map[string]any{
		"ok":      true,
		"message": "Код отправлен",
	})
}

func (h *OtpHandler) VerifySmsOtp(w http.ResponseWriter, r *http.Request) {
	log := h.log.With("method", "VerifySmsOtp")
	log.Info("request_received")

	var req struct {
		PhoneNumber string `json:"phone_number"`
		Code        string `json:"code"`
	}

	if err := json.ParseJSON(r, &req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.otpService.VerifySmsOtp(r.Context(), &userv1.VerifySmsOtpRequest{
		PhoneNumber: req.PhoneNumber,
		Code:        req.Code,
	})
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusUnauthorized, errors.New("verification failed"))
		return
	}

	log.Info("succeeded", "user_id", resp.GetUser().GetId())
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OtpHandler) RequestMailOtp(w http.ResponseWriter, r *http.Request) {
	log := h.log.With("method", "RequestMailOtp")
	log.Info("request_received")

	var req userv1.RequestMailOtpRequest
	if err := json.ParseProtoJSON(r.Body, &req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.otpService.RequestMailOtp(r.Context(), &req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusInternalServerError, errors.New("failed to request mail OTP"))
		return
	}

	log.Info("succeeded", "email", req.Email)
	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OtpHandler) VerifyMailOtp(w http.ResponseWriter, r *http.Request) {
	log := h.log.With("method", "VerifyMailOtp")
	log.Info("request_received")

	var req userv1.VerifyMailOtpRequest
	if err := json.ParseProtoJSON(r.Body, &req); err != nil {
		log.Error("invalid_payload", "error", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.otpService.VerifyMailOtp(r.Context(), &req)
	if err != nil {
		log.Error("gRPC_call_failed", "error", err)
		json.WriteError(w, http.StatusUnauthorized, errors.New("verification failed"))
		return
	}

	log.Info("succeeded", "user_id", resp.GetUser().GetId())
	json.WriteJSON(w, http.StatusOK, resp)
}
