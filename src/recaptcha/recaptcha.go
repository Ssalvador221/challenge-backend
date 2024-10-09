package recaptcha

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
)

type RecaptchaResponse struct {
	Success bool    `json:"success"`
	Score   float64 `json:"score"`
}

func ValidateCaptcha(recaptchaResponse string) error {
	secretKey := os.Getenv("RECAPTCHA_SECRET_KEY")
	recaptchaURL := "https://www.google.com/recaptcha/api/siteverify"

	data := url.Values{}
	data.Set("secret", secretKey)
	data.Set("response", recaptchaResponse)

	resp, err := http.PostForm(recaptchaURL, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result RecaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if !result.Success {
		return errors.New("reCAPTCHA validation failed")
	}

	return nil
}
