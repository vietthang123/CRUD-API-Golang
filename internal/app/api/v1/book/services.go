package book

type IBookServices interface {
	CreateBook(formData Book) (Book, error)
	BulkInsert(data []Book) ([]Book, error)
	GetListBook(limit int, offSet int, order string, filter map[string]interface{}) ([]Book, int, error)
	GetBookByID(bookId int) (Book, error)
	UpdateBook(bookId int, formData Book) (Book, error)
	DeleteBook(bookId int) (Book, error)
}

type bookServices struct {
	bookRepo IBookRepository
}

func NewBookServices(repository IBookRepository) *bookServices {
	return &bookServices{bookRepo: repository}
}

func (b *bookServices) CreateBook(formData Book) (Book, error) {
	return b.bookRepo.CreateBook(formData)
}

func (b *bookServices) BulkInsert(data []Book) ([]Book, error) {
	return b.bookRepo.BulkInsert(data)
}

func (b *bookServices) GetListBook(limit int, offSet int, order string, filter map[string]interface{}) ([]Book, int, error) {
	return b.bookRepo.GetListBook(limit, offSet, order, filter)
}

func (b *bookServices) GetBookByID(bookId int) (Book, error) {
	return b.bookRepo.GetBookByID(bookId)
}

func (b *bookServices) UpdateBook(bookId int, formData Book) (Book, error) {
	return b.bookRepo.UpdateBook(bookId, formData)
}

func (b *bookServices) DeleteBook(bookId int) (Book, error) {
	return b.bookRepo.DeleteBook(bookId)
}
