//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/snoeffels/mygo/controllers"
	"github.com/snoeffels/mygo/repositories"
	"github.com/snoeffels/mygo/services"
	"gorm.io/gorm"
)

func initTodoAPI(db *gorm.DB) controllers.TodoAPI {
	wire.Build(
		repositories.ProvideTodoRepository,
		services.ProvideTodoService,
		controllers.ProvideTodoAPI)

	return controllers.TodoAPI{}
}
