//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/snoeffels/mygo/api"
	"github.com/snoeffels/mygo/persistence"
	"gorm.io/gorm"
)

func initTodoAPI(db *gorm.DB) api.TodoController {
	wire.Build(
		persistence.ProvideTodoRepository,
		persistence.ProvideTodoService,
		api.ProvideTodoAPI)

	return api.TodoController{}
}
