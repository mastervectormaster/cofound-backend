package handler

import (
	"github.com/mastervectormaster/cofound-backend/user"
)

type Handler struct {
	userStore user.Store
	validator *Validator
}

func NewHandler(us user.Store) *Handler {
	v := NewValidator()
	return &Handler{
		userStore: us,
		validator: v,
	}
}
