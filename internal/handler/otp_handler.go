package handler

import (
	"log/slog"

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

// func (h *OtpHandler) RequestOtp(w http.ResponseWriter, r *http.Request) {
// 	var req struct {
// 		PhoneNumber string `json:"phone_number"`
// 	}

// 	if err := json.ParseJSON(r, &req); err != nil {
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	_, err := h.otpService.RequestOtp(r.Context(), &userv1.RequestOtpRequest{
// 		PhoneNumber: req.PhoneNumber,
// 	})
// 	if err != nil {
// 		h.log.Error("failed to request OTP", "err", err)
// 		json.WriteError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	json.WriteJSON(w, http.StatusOK, map[string]any{
// 		"ok":      true,
// 		"message": "Код отправлен",
// 	})
// }

// func (h *OtpHandler) VerifyOtp(w http.ResponseWriter, r *http.Request) {
// 	var req struct {
// 		PhoneNumber string `json:"phone_number"`
// 		Code        string `json:"code"`
// 	}

// 	if err := json.ParseJSON(r, &req); err != nil {
// 		json.WriteError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	resp, err := h.otpService.VerifyOtp(r.Context(), &userv1.VerifyOtpRequest{
// 		PhoneNumber: req.PhoneNumber,
// 		Code:        req.Code,
// 	})
// 	if err != nil {
// 		h.log.Error("failed to verify OTP", "err", err)
// 		json.WriteError(w, http.StatusUnauthorized, err)
// 		return
// 	}

// 	json.WriteJSON(w, http.StatusOK, resp)
// }
