package repository

//go:generate mockgen -source=interfaces.go -destination=mocks/repository_mock.go -package=repository_mock

type Repository interface {
	CreateBook() *Book
	GetAllBooks() []Book
	GetBookById(Id int) *Book
	DeleteBook(Id int) Book
}
