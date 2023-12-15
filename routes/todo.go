package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/snoeffels/mygo/controllers"
	"github.com/snoeffels/mygo/middlewares"
)

func TodoRoute(r *gin.Engine, todo controllers.TodoAPI) *gin.Engine {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1").Use(middlewares.TokenAuthMiddleware())
		{
			v1.GET("/todo", todo.FindAll)
			v1.POST("/todo", todo.Create)
			v1.GET("/todo/:id", todo.FindById)
			v1.PUT("/todo/:id", todo.Update)
			v1.DELETE("/todo/:id", todo.Delete)
		}
	}

	return r
}
