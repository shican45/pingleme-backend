//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// EvaluationTable 评审表模型
type EvaluationTable struct {
	gorm.Model
	HomeworkID int `gorm:"type:int;not null"`
	TeamID     int `gorm:"type:int;not null"`
}

// EvaluationTableItem 评审表项模型
type EvaluationTableItem struct {
	gorm.Model
	EvaluationTableID int    `gorm:"type:int;not null"`
	Content           string `gorm:"type:varchar(255);not null"`
	Score             int    `gorm:"type:int;not null"`
	Suggest           string `gorm:"type:text"`
	ParentItemID      int    `gorm:"type:int;not null"`
	Sequence          int    `gorm:"type:int;not null"`
}
