<h3>
Bookstore REST API
</h3>

An API with ability to create, read, update, delete books in db.
<h3>
Stack
</h3>

<ul>
<li>
Go 1.18
</li>
<li>
MySQL 8.0
</li>
<li>
Docker
</li>dfsdfsdf

<h3>
Libraries
</h3>

All dependencies you can see in `go.mod`


`github.com/go-sql-driver/mysql v1.5.0`

`github.com/gorilla/mux v1.8.0`

`github.com/jinzhu/gorm v1.9.16`

`github.com/jinzhu/inflection v1.0.0`

`github.com/joho/godotenv v1.4.0`

<h3>Postman</h3>

All requests and routes was tested with Postman

`/book` use "POST" method to create new book

`/book` use "GET" method to get all books in JSON

`/book/{bookId}` use "GET" method to get book by bookId in JSON

`/book/{bookId}` use "PUT" method to update book by bookId

`/book/{bookId}` use "DELETE" method to delete book by bookId

<h3>Launch</h3>

Use `git clone https://github.com/rogaliiik/go-bookstore.git`

Install dependecies with `go mod download`

Compile project with `go build -o bookstore.exe ./cmd`

Start server with `./bookstore`


