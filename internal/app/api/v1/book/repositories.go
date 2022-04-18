package book

import (
	"github.com/go-pg/pg/v10"
)

type IBookRepository interface {
	CreateBook(formData Book) (Book, error)
	BulkInsert(data []Book) ([]Book, error)
	GetListBook(limit int, offSet int, order string, filter map[string]interface{}) ([]Book, int, error)
	GetBookByID(bookId int) (Book, error)
	UpdateBook(bookId int, formData Book) (Book, error)
	DeleteBook(bookId int) (Book, error)
}

type bookRepository struct {
	dbpg *pg.DB
}

func NewBookRepository(dbpg *pg.DB) *bookRepository {
	return &bookRepository{dbpg: dbpg}
}

func (b *bookRepository) CreateBook(formData Book) (Book, error) {
	_, err := b.dbpg.Model(&formData).Insert()
	if err != nil {
		return Book{}, err
	}
	return formData, nil
}

func (b *bookRepository) BulkInsert(data []Book) ([]Book, error) {
	_, err := b.dbpg.Model(&data).Insert()
	if err != nil {
		return []Book{}, err
	}
	return data, nil
}

func (b *bookRepository) GetBookByID(bookId int) (Book, error) {
	book := &Book{}
	err := b.dbpg.Model(book).Where("id = ?", bookId).Select()
	if err != nil {
		return Book{}, err
	}
	return *book, nil
}

func (b *bookRepository) GetListBook(limit int, offSet int, order string, filter map[string]interface{}) ([]Book, int, error) {
	var book []Book
	var query *pg.Query
	query = b.dbpg.Model(&book)
	if limit != 0 {
		query.Limit(limit)
	}

	if offSet == 0 {
		query.Offset(10)
	} else if offSet != 0 {
		query.Offset(offSet)
	}

	if order != "" {
		query.Order(order)
	}
	for object, v := range filter {
		query = query.Where("? = ?", pg.Ident(object), v)
	}
	count, _ := query.Count()
	err := query.Select()
	if err != nil {
		return []Book{}, count, err
	}
	return book, count, nil
}

func (b *bookRepository) UpdateBook(bookId int, formData Book) (Book, error) {
	book := &Book{}
	book.Name = formData.Name
	book.Category = formData.Category
	book.Content = formData.Content
	book.Author = formData.Author
	_, err := b.dbpg.Model(book).Where("id = ?", bookId).Update()
	if err != nil {
		return Book{}, err
	}
	return *book, nil
}

func (b *bookRepository) DeleteBook(bookId int) (Book, error) {
	book := &Book{}
	query := b.dbpg.Model(book).Where("id = ? ", bookId)
	err := query.Select()
	if err != nil {
		return Book{}, err
	}
	query.Delete()
	return *book, nil
}
