package validate

import (
	"log/slog"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Initialize validator
var V *validator.Validate

func init() {
	V = validator.New()
	register(V)
}

func validateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasUpper {
		slog.Info("password don't have uppercase")
	}
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasDigit {
		slog.Info("password don't have digits")
	}

	return hasUpper && hasDigit
}

func register(V *validator.Validate) {
	V.RegisterValidation("passwd", validateStrongPassword)
}
