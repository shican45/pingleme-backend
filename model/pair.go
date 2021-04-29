//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Pair 结对模型
type Pair struct {
	gorm.Model
	Student1ID int `gorm:"type:int;not null;unique_index:studentID"`
	Student2ID int `gorm:"type:int;unique_index:studentID"`
}

// GetPair 用ID获取结对
func (Repo *Repository) GetPair(ID int) (Pair, error){
	var pair Pair
	result := Repo.DB.First(&pair, ID)
	if result.Error != nil {
		return Pair{}, result.Error
	}
	return pair, nil
}

// CreatePair 创建结对
func (Repo *Repository) CreatePair(pair Pair) (Pair, error){
	result := Repo.DB.Create(&pair)
	if result.Error != nil {
		return Pair{}, result.Error
	}
	return pair, nil
}

// GetPairByStudentID 根据学生ID获取结对
func (Repo *Repository) GetPairByStudentID(ID int) (Pair, error){
	var pair Pair
	result := Repo.DB.Where("Student1ID = ?", ID).Or("Student2ID = ?", ID).First(&pair)
	if result.Error != nil {
		return Pair{}, result.Error
	}
	return pair, nil
}

// DeletePair 根据ID删除结对
func (Repo *Repository) DeletePair(ID int) error{
	result := Repo.DB.Delete(&Pair{}, ID)
	return result.Error
}

// DeletePairByStudentID 根据学生ID删除结对
func (Repo *Repository) DeletePairByStudentID(ID int) error{
	result := Repo.DB.Where("Student1ID = ?", ID).Or("Student2ID = ?", ID).Delete(&Pair{})
	return result.Error
}

// UpdatePair 更新Pair
func (Repo *Repository) UpdatePair(ID int, student1ID int, student2ID int) (Pair, error){
	var pair Pair
	result := Repo.DB.First(&pair, ID)
	if result.Error != nil {
		return Pair{}, result.Error
	}

	if student1ID != 0 {
		pair.Student1ID = student1ID
	}
	if student2ID != 0 {
		pair.Student2ID = student2ID
	}
	Repo.DB.Save(&pair)
	return pair, nil
}