package handlers

import (
	"backend-challenge/src/mailer"
	"backend-challenge/src/recaptcha"
	"encoding/json"
	"net/http"
	"net/mail"
)

type Handlers struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Comment           string `json:"comment"`
	RecaptchaResponse string `json:"g-recaptcha-response"`
}

type ErrorResponse struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func HandleContactForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{
			Type:     "about:blank",
			Title:    "MethodNotAllowed",
			Detail:   "Only POST method is allowed",
			Instance: r.URL.Path,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var handler Handlers

	if err := json.NewDecoder(r.Body).Decode(&handler); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Type:     "about:blank",
			Title:    "BadRequestError",
			Detail:   "Invalid request payload",
			Instance: r.URL.Path,
		})
		return
	}

	if _, err := mail.ParseAddress(handler.Email); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Type:     "about:blank",
			Title:    "BadRequestError",
			Detail:   "The email is invalid",
			Instance: r.URL.Path,
		})
		return
	}

	err := recaptcha.ValidateCaptcha(handler.RecaptchaResponse)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{
			Type:     "about:blank",
			Title:    "UnauthorizedError",
			Detail:   "The captcha is incorrect!",
			Instance: r.URL.Path,
		})
		return
	}

	if err := mailer.SendContactEmail(handler.Name, handler.Email, handler.Comment); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{
			Type:     "about:blank",
			Title:    "InternalServerError",
			Detail:   "Failed to send email.",
			Instance: r.URL.Path,
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
