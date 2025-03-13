package handlers

import (
	"imobiliaria/internal/usecases"

	"github.com/go-playground/validator/v10"
)

// Implements implicito
type Handler struct {
	Usecases  usecases.Usecases
	Validator *validator.Validate
}
