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

// GetAllTeachers 获得该班级的所有老师
func (class *Class) GetAllTeachers() ([]User, error){
	var teachers []User
	teachers = class.Teachers
	return teachers, nil
}

// GetAllStudents 获得该班级的所有学生
func (class *Class) GetAllStudents() ([]User, error){
	var students []User
	students = class.Students
	return students, nil
}

// AddTeacher 添加一个老师
func (class *Class) AddTeacher(teacher User) error{
	var classID = class.ID
	var teacherID = teacher.ID
	result := Repo.DB.Exec("insert into teacher_class(class_id,teacher_id) values(?,?)", classID, teacherID)
	return result.Error
}

// AddClass 添加一个班级
func (Repo *Repository) AddClass(Name string) error {
	result := Repo.DB.Exec("insert into class(name) values(?)", Name)
	return result.Error
}

// AddStudent 添加一个学生
func (class *Class) AddStudent(student User) error{
	var classID = class.ID
	var studentID = student.ID
	result := Repo.DB.Exec("insert into teacher_class(class_id,teacher_id) values(?,?)", classID, studentID)
	return result.Error
}

// DeleteClass 删除班级
func (class *Class) DeleteClass() error{
	result := Repo.DB.Exec("delete from class where class_id = ?", class.ID)
	return result.Error
}

// DeleteTeacher 删除一个老师
func (class *Class) DeleteTeacher(teacher User) error {
	var classID = class.ID
	var teacherID = teacher.ID
	result := Repo.DB.Exec("delete from teacher_class where class_id = ? and teacher_id = ?", classID, teacherID)
	return result.Error
}

// DeleteStudent 删除一个学生
func (class *Class) DeleteStudent(student User) error {
	var classID = class.ID
	var studentID = student.ID
	result := Repo.DB.Exec("delete from student_class where class_id = ? and student_id = ?", classID, studentID)
	return result.Error
}

// UpdateClassName 修改班级名字
func (class *Class) UpdateClassName(Name string) error {
	result := Repo.DB.Exec("update class set name = ?", Name)
	return result.Error
}