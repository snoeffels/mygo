# Golang-Gin-Gorm-Postgres boilerplate

This is a template project for upcoming golang project ideas.
It uses gorm as ORM with Postgres and also uses JWT auth middleware 
if GIN_MODE=release is set in .env / .env-docker

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
 swag init -parseDependency=true
```

### Generate dependency injection using googles go:wire
```shell
go run github.com/google/wire/cmd/wire
```