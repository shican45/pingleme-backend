//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Team 团队模型
type Team struct {
	gorm.Model
	Number        int    `gorm:"type:int;not null"`
	Name          string `gorm:"type:varchar(255);not null;unique"`
	GroupLeaderID int    `gorm:"type:int;not null"`
	ClassID       int    `gorm:"type:int;not null"`
	Students      []User `gorm:"many2many:student_team;"`
}

// GetTeamByNumAndClassID
func (Repo *Repository) GetTeamByNumAndClassID(Num interface{}, ClassID interface{}) (Team, error){
	var team Team
	result := Repo.DB.Where("number = ? and class_id = ?", Num, ClassID).First(&team)
	return team, result.Error
}

// GetAllStudents 获得本团队所有的学生
func (team *Team) GetAllStudents() ([]User, error){
	var students []User
	students = team.Students
	return students, nil
}