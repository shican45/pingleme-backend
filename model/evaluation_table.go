//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// EvaluationTable 评审表模型
type EvaluationTable struct {
	gorm.Model
	TableName  string `gorm:"type:varchar(255);not null"`
	HomeworkID int    `gorm:"type:int;not null"`
	TeamID     int    `gorm:"type:int;not null"`
	TableItems []EvaluationTableItem
}

// EvaluationTableItem 评审表项模型
type EvaluationTableItem struct {
	gorm.Model
	EvaluationTableID uint
	Content           string `gorm:"type:varchar(255);not null"`
	Score             int    `gorm:"type:int;not null;default:-1"`
	Description       string `gorm:"type:text"`
	ParentItemID      int    `gorm:"type:int;not null;default:0"`
	Index             int    `gorm:"type:int;not null;default:0"`
}

// GetEvaluationTable 获取评审表
func (Repo *Repository) GetEvaluationTable(ID uint) (EvaluationTable, error) {
	var table EvaluationTable
	tableResult := Repo.DB.First(&table, ID)
	if tableResult.Error != nil {
		return EvaluationTable{}, tableResult.Error
	}
	err := Repo.DB.Model(&table).Association("TableItems").Find(&table.TableItems)
	if err != nil {
		return EvaluationTable{}, err
	}

	return table, nil
}

// SetEvaluationTable 保存评审表
func (Repo *Repository) SetEvaluationTable(table EvaluationTable) error {
	result := Repo.DB.Create(&table)
	return result.Error
}
