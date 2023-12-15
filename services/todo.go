package services

import (
	"github.com/snoeffels/mygo/models"
	"github.com/snoeffels/mygo/repositories"
)

type TodoService struct {
	TodoRepository repositories.TodoRepository
}

func ProvideTodoService(t repositories.TodoRepository) TodoService {
	return TodoService{TodoRepository: t}
}

func (t *TodoService) FindAll() ([]models.Todo, error) {
	return t.TodoRepository.FindAll()
}

func (t *TodoService) FindById(id uint) (models.Todo, error) {
	return t.TodoRepository.FindById(id)
}

func (t *TodoService) Create(todo *models.Todo) error {
	err := t.TodoRepository.Create(todo)
	return err
}

func (t *TodoService) Update(todo, updateTodo *models.Todo) error {
	err := t.TodoRepository.Update(todo, updateTodo)
	return err
}

func (t *TodoService) DeleteById(todo *models.Todo, id uint) error {
	err := t.TodoRepository.DeleteById(todo, id)
	return err
}
