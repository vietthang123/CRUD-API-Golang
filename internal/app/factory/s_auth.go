package factory

import (
	"developer-orientenergy-golang/internal/app/api/v1/auth"
)

func GetAuthService() auth.IAuthService {
	repo := GetAuthRepository()
	return auth.NewAuthService(repo)
}
