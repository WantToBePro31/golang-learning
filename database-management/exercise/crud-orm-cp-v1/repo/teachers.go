package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	if err := t.db.Create(&data).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	var teachers []model.Teacher
	if err := t.db.Unscoped().Find(&teachers).Error; err != nil {
		return []model.Teacher{}, err
	}
	return teachers, nil // TODO: replace this
}

func (t TeacherRepo) Update(id uint, name string) error {
	var teacher model.Teacher
	if err := t.db.Model(&teacher).Where("id = ?", id).Update("name", name).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (t TeacherRepo) Delete(id uint) error {
	var teacher model.Teacher
	if err := t.db.Where("id = ?", id).Delete(&teacher).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}
