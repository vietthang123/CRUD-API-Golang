package factory

import (
	"developer-orientenergy-golang/internal/app/api/v1/school"
	"developer-orientenergy-golang/internal/pkg/database"
)

func GetSchoolRepository() school.ISchoolRepository {
	r := database.GetConnection()
	return school.NewSchoolRepository(r.Db)
}
