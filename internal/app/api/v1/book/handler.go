package book

import (
	"developer-orientenergy-golang/internal/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type bookController struct {
	bookService IBookServices
}

func NewBookController(bookService IBookServices) *bookController {
	return &bookController{bookService: bookService}
}

func (b *bookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var formData Book
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &formData)
	if err != nil {
		util.RespondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	book, err := b.bookService.CreateBook(formData)
	if err != nil {
		util.RespondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.RespondJSONSuccess(w, fmt.Sprintf("Create Book %d Success", book.ID))
	return
}

func (b *bookController) BulkInsert(w http.ResponseWriter, r *http.Request) {
	var data []Book
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &data)
	_, err := b.bookService.BulkInsert(data)
	if err != nil {
		util.RespondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.RespondJSONSuccess(w, fmt.Sprintf("Insert Data Success"))
	return
}

func (b *bookController) GetListBook(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	paramsRequest := util.ParseToParam(query, r)

	limitRequest := paramsRequest.Limit
	pageRequest := paramsRequest.Page
	order := paramsRequest.Order
	filter := paramsRequest.FilterObject

	paginationRequest := util.NewPaginationRequest(limitRequest, pageRequest)
	limit := paginationRequest.Limit
	offSet := paginationRequest.Offset

	book, count, _ := b.bookService.GetListBook(limit, offSet, order, filter)
	util.RespondJSON(w, http.StatusOK, util.PaginationResponse(paginationRequest, count, book))
	return
}

func (b *bookController) GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	book, err := b.bookService.GetBookByID(id)
	if err != nil {
		util.RespondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.RespondJSONSuccess(w, book)
	return
}

func (b *bookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var formData Book
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &formData)
	if err != nil {
		util.RespondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	book, err := b.bookService.UpdateBook(id, formData)
	if err != nil {
		util.RespondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.RespondJSONSuccess(w, book)
	return
}

func (b *bookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	book, err := b.bookService.DeleteBook(id)
	if err != nil {
		util.RespondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.RespondJSONSuccess(w, fmt.Sprintf("Delete Book %d Success", book.ID))
	return
}
