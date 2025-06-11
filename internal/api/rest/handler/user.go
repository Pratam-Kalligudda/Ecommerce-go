package handler

import (
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/api/rest"
	"github.com/Pratam-Kalligudda/Ecommerce-go/internal/service"
)

type UserHandler struct {
	svc service.UserService
}

func SetupUserRoutes(rh *rest.Handler) {
	// app := rh.App
	// svc := service.UserService{Repo: repository.NewRepository(rh.DB)}
	// handler := handler.UserHandler{svc: svc}
}
