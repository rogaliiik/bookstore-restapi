package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/rogaliiik/bookstore/pkg/config"
	"github.com/rogaliiik/bookstore/pkg/utils"
)

var (
	Repo BookRepository
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) CreateBook(b *Book) *Book {
	r.DB.NewRecord(b)
	r.DB.Create(&b)
	return b
}

func (r *BookRepository) GetAllBooks() []Book {
	var Books []Book
	r.DB.Find(&Books)
	return Books
}

func (r *BookRepository) GetBookById(Id int) *Book {
	var getBook Book
	r.DB.Where("ID=?", Id).Find(&getBook)
	return &getBook
}

func (r *BookRepository) DeleteBook(Id int) Book {
	var book Book
	r.DB.Where("ID=?", Id).Delete(book)
	return book
}

func Connect() {
	dbPassword := utils.GetEnvVar("POSTGRES_PASSWORD")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Conf.DB.Host,
		config.Conf.DB.Port,
		config.Conf.DB.User,
		dbPassword,
		config.Conf.DB.DBName,
		config.Conf.DB.Sslmode,
	)

	database, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	Repo.DB = database
}

func init() {
	Connect()
	Repo.DB.AutoMigrate(&Book{})
}
