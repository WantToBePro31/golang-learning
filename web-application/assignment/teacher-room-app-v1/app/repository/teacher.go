package handler

import (
	"a21hc3NpZ25tZW50/app/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (u *TeacherRepo) AddTeacher(teacher model.Teacher) error {
	if err := u.db.Create(&teacher).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *TeacherRepo) ReadTeacher() ([]model.ViewTeacher, error) {
	var viewTeacher []model.ViewTeacher	
	if err := u.db.Table("teachers").Select("name, field_of_study, age").Where("deleted_at IS NULL").Find(&viewTeacher).Error; err != nil {
		return []model.ViewTeacher{}, err 
	}
	return viewTeacher, nil // TODO: replace this
}

func (u *TeacherRepo) UpdateName(id uint, name string) error {
	var teacher model.Teacher
	if err := u.db.Model(&teacher).Where("id = ?", id).Update("name", name).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *TeacherRepo) DeleteTeacher(id uint) error {
	var teacher model.Teacher
	if err := u.db.Where("id = ?", id).Delete(&teacher).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}
