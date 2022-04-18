package factory

import (
	"developer-orientenergy-golang/internal/app/api/v1/book"
)

func GetBookService() book.IBookServices {
	repo := GetBookRepository()
	return book.NewBookServices(repo)
}
