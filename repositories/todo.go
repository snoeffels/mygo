package repositories

import (
	"github.com/snoeffels/mygo/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func ProvideTodoRepository(DB *gorm.DB) TodoRepository {
	return TodoRepository{DB: DB}
}

func (t *TodoRepository) FindAll() (todo []models.Todo, err error) {
	return todo, t.DB.Find(&todo).Error
}

func (t *TodoRepository) FindById(id uint) (todo models.Todo, err error) {
	return todo, t.DB.First(&todo, id).Error
}

func (t *TodoRepository) Create(todo *models.Todo) error {
	if err := t.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (t *TodoRepository) Update(todo, updateTodo *models.Todo) error {
	if err := t.DB.Model(&todo).Updates(updateTodo).Error; err != nil {
		return err
	}
	return nil
}

func (t *TodoRepository) DeleteById(todo *models.Todo, id uint) error {
	if err := t.DB.Where("id = ?", id).Delete(todo).Error; err != nil {
		return err
	}
	return nil
}
