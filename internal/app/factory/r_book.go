package factory

import (
	"developer-orientenergy-golang/internal/app/api/v1/book"
	"developer-orientenergy-golang/internal/pkg/database"
)

func GetBookRepository() book.IBookRepository {
	r := database.GetConnection()
	return book.NewBookRepository(r.PgDb)
}
