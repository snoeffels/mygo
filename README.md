# Golang-Gin-Gorm Project Template

Go Gin WebFramework + Gorm + Wire

- [Gin WebFramework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/index.html)
- [Wire](https://github.com/google/wire)

```shell
Controllers -> Services -> Repositories -> Models 
```

```shell
Dependency Injection
wire

'go run github.com/google/wire/cmd/wire'
```

```shell
go run .
```

```shell
go build .
./main
```

After modify swagger options
```shell
 swag init -parseDependency=true
```