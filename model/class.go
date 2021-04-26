//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Class 班级模型
type Class struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null;unique"`
	Teachers []User `gorm:"many2many:teacher_class;"`
	Students []User `gorm:"many2many:student_class;"`
}

// GetClassByID 通过班级ID获取班级
func (Repo *Repository) GetClassByID(ID interface{}) (Class, error){
	var class Class
	result := Repo.DB.First(&class, ID)
	return class, result.Error
}
