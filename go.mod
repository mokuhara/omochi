module omochi

// +heroku goVersion 1.14
// +heroku install -tags 'mysql' ./vendor/github.com/golang-migrate/migrate/v4/cmd/migrate

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.0.0-20201117144127-c1f2f97bffc9
)
