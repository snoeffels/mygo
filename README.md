# Golang-Gin-Gorm-Postgres boilerplate

This is a template project for upcoming golang project ideas.
It uses gorm as ORM with Postgres and also uses JWT auth middleware 
if GIN_MODE=release is set in .env / .env-docker

The project also offers a generated swagger ui at [http://localhost:8080/swagger-ui/index.html](http://localhost:8080/swagger-ui/index.html)
- eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.vZa-SKJgVgooNmaQeYbk8KxbsSohjoDm0-_rxm96FZw

### Resources
- [Gin WebFramework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/index.html)
- [Wire](https://github.com/google/wire)

### Run using docker compose
```shell
docker compose up --build -d
```

### Run local
```shell
go run .
```

### Build & run local
```shell
go build .
./main
```

### Test curl
```shell
curl https://reqbin.com/echo/get/json
   -H "Accept: application/json"
   -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.vZa-SKJgVgooNmaQeYbk8KxbsSohjoDm0-_rxm96FZw"
```

### Generate swagger
```shell
 go install github.com/swaggo/swag/cmd/swag@latest
 
 swag init -parseDependency=true
```

### Generate dependency injection using googles go:wire
```shell
go install github.com/google/wire/cmd/wire

go run github.com/google/wire/cmd/wire
```