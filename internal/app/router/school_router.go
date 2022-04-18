package router

import (
	"developer-orientenergy-golang/internal/app/api/v1/school"
	"developer-orientenergy-golang/internal/app/factory"
	"net/http"
)

func SchoolRouter(Dispatch map[string]http.HandlerFunc) map[string]http.HandlerFunc {
	schoolService := factory.GetSchoolService()
	authController := school.NewSchoolController(schoolService)
	Dispatch["CreateSchool"] = authController.CreateSchool
	return Dispatch
}
