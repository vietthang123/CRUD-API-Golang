package router

import (
	"developer-orientenergy-golang/internal/app/api/v1/auth"
	"developer-orientenergy-golang/internal/app/factory"
	"net/http"
)

func AuthRouter(Dispatch map[string]http.HandlerFunc) map[string]http.HandlerFunc {
	authService := factory.GetAuthService()
	authController := auth.NewAuthController(authService)
	Dispatch["CheckLogin"] = authController.CheckLogin
	Dispatch["InsertUser"] = authController.InsertUser
	return Dispatch
}
