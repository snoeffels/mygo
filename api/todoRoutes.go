package api

import (
	"github.com/gin-gonic/gin"
	"github.com/snoeffels/mygo/middleware"
)

func TodoRoute(r *gin.Engine, todo TodoAPI) *gin.Engine {
	api := r.Group("/api")
	{
		v1 := api.Group("/v1").Use(middleware.TokenAuthMiddleware())
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
