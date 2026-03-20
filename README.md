## Init ##
go mod init golang-playground
go run cmd/api/main.go
go mod build
go get "url"
go mod vendor

## Environment ##
gin
gorm
mysql / psql
godotenv
bcrypt
migrate
validator
zap
redis
