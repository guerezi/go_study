package handlers

import (
	"imobiliaria/internal/usecases"
	"imobiliaria/internal/validator"
)

type Handler struct {
	usecases  usecases.Usecases    // sem ponteiro pq é uma interface
	validator *validator.Validator // ponteiro porque é um struct
}

func NewHandler(usecases usecases.Usecases, validator *validator.Validator) *Handler {
	return &Handler{
		usecases:  usecases,
		validator: validator,
	}
}
