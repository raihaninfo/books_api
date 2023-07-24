# books_api

```bash
require (
	github.com/gin-gonic/gin v1.9.1
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/google/uuid v1.3.0
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.17
	golang.org/x/crypto v0.11.0
	gorm.io/driver/sqlite v1.5.2
	gorm.io/gorm v1.25.2
)

```

## Public api routes

```bash
[Register new user]        POST     /api/register
[Login with Jwt token]     POST     /api/login

```

## Protected api routes

```bash
[Insert Book]              POST     /api/book
[Get All Book]             GET      /api/books
[Get Single Book]          GET      /api/book/:id
[Update single book]       PUT      /api/book/:id
[Delete book]              DELETE   /api/book/:id

```
