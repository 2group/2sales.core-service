package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	userv1 "github.com/2group/2sales.core-service/pkg/gen/go/user"
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

func (h *OtpHandler) RequestOtp(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PhoneNumber string `json:"phone_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	_, err := h.otpService.RequestOtp(r.Context(), &userv1.RequestOtpRequest{
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		h.log.Error("failed to request OTP", "err", err)
		http.Error(w, "не удалось отправить код", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"ok":      true,
		"message": "Код отправлен",
	})
}

func (h *OtpHandler) VerifyOtp(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PhoneNumber string `json:"phone_number"`
		Code        string `json:"code"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.otpService.VerifyOtp(r.Context(), &userv1.VerifyOtpRequest{
		PhoneNumber: req.PhoneNumber,
		Code:        req.Code,
	})
	if err != nil {
		h.log.Error("failed to verify OTP", "err", err)
		http.Error(w, "неверный или просроченный код", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
