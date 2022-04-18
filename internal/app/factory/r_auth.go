package factory

import (
	"developer-orientenergy-golang/internal/app/api/v1/auth"
	"developer-orientenergy-golang/internal/pkg/database"
)

func GetAuthRepository() auth.IAuthRepository {
	r := database.GetConnection()
	return auth.NewAuthRepository(r.Db, r.PgDb)
}
