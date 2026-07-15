module taskflow-backend

go 1.22.2

replace golang.org/x/crypto => github.com/golang/crypto v0.23.0

replace golang.org/x/sys => github.com/golang/sys v0.20.0

require (
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/google/uuid v1.6.0
	github.com/mattn/go-sqlite3 v1.14.48
	golang.org/x/crypto v0.23.0
)

require github.com/joho/godotenv v1.5.1
