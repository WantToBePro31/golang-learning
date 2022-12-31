package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return TodoRepository{db}
}

func (u *TodoRepository) AddTodo(todo model.Todo) error {
	if err := u.db.Create(&todo).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *TodoRepository) ReadTodo() ([]model.Todo, error) {
	var todos []model.Todo
	if err := u.db.Where("deleted_at IS NULL").Find(&todos).Error; err != nil {
		return []model.Todo{}, err
	}
	return todos, nil // TODO: replace this
}

func (u *TodoRepository) UpdateDone(id uint, status bool) error {
	var todo model.Todo
	if err := u.db.Model(&todo).Where("id = ?", id).Update("done", status).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *TodoRepository) DeleteTodo(id uint) error {
	var todo model.Todo
	if err := u.db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}
