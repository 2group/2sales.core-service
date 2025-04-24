package handler

import (
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
	var req struct {
		PhoneNumber string `json:"phone_number"`
	}

	if err := json.ParseJSON(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.otpService.RequestSmsOtp(r.Context(), &userv1.RequestSmsOtpRequest{
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		h.log.Error("failed to request OTP", "err", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, map[string]any{
		"ok":      true,
		"message": "Код отправлен",
	})
}

func (h *OtpHandler) VerifySmsOtp(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PhoneNumber string `json:"phone_number"`
		Code        string `json:"code"`
	}

	if err := json.ParseJSON(r, &req); err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.otpService.VerifySmsOtp(r.Context(), &userv1.VerifySmsOtpRequest{
		PhoneNumber: req.PhoneNumber,
		Code:        req.Code,
	})
	if err != nil {
		h.log.Error("failed to verify OTP", "err", err)
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OtpHandler) RequestMailOtp(w http.ResponseWriter, r *http.Request) {
	var req userv1.RequestMailOtpRequest
	if err := json.ParseProtoJSON(r.Body, &req); err != nil {
		h.log.Error("failed to parse request JSON", "err", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.otpService.RequestMailOtp(r.Context(), &req)
	if err != nil {
		h.log.Error("failed to request mail OTP", "err", err)
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, resp)
}

func (h *OtpHandler) VerifyMailOtp(w http.ResponseWriter, r *http.Request) {
	var req userv1.VerifyMailOtpRequest
	if err := json.ParseProtoJSON(r.Body, &req); err != nil {
		h.log.Error("failed to parse request JSON", "err", err)
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := h.otpService.VerifyMailOtp(r.Context(), &req)
	if err != nil {
		h.log.Error("failed to verify mail OTP", "err", err)
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	json.WriteJSON(w, http.StatusOK, resp)
}
