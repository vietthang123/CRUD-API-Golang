package router

import (
	"developer-orientenergy-golang/internal/app/api/v1/book"
	"developer-orientenergy-golang/internal/app/factory"
	"net/http"
)

func BookRouter(Dispatch map[string]http.HandlerFunc) map[string]http.HandlerFunc {
	bookService := factory.GetBookService()
	bookController := book.NewBookController(bookService)
	Dispatch["CreateBook"] = bookController.CreateBook
	Dispatch["BulkInsert"] = bookController.BulkInsert
	Dispatch["GetListBook"] = bookController.GetListBook
	Dispatch["GetBookByID"] = bookController.GetBookByID
	Dispatch["UpdateBook"] = bookController.UpdateBook
	Dispatch["DeleteBook"] = bookController.DeleteBook
	return Dispatch
}
