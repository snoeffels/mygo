package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/snoeffels/mygo/config"
	_ "github.com/snoeffels/mygo/docs"
	"github.com/snoeffels/mygo/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /api/v1
// @query.collection.format multi

// @title Todo API
// @description Todo microservice API
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

type dbHandlers struct {
	db *gorm.DB
}

func main() {
	var err error
	var dbHandler dbHandlers

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHandler.db, err = config.InitPostgres()
	if err != nil {
		panic(err)
	}

	sqlDB, _ := dbHandler.db.DB()

	defer func(sqlDB *sql.DB) {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Could not close db connection")
		}
	}(sqlDB)

	r := initRoutes(dbHandler)
	log.Fatal(r.Run())
}

func initRoutes(dbHandler dbHandlers) *gin.Engine {
	r := gin.Default()

	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	r.Use(cors.New(conf))

	todoAPI := initTodoAPI(dbHandler.db)
	routes.TodoRoute(r, todoAPI)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
