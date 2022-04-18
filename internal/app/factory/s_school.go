package factory

import (
	"developer-orientenergy-golang/internal/app/api/v1/school"
)

func GetSchoolService() school.ISchoolServices {
	repo := GetSchoolRepository()
	return school.NewSchoolServices(repo)
}
