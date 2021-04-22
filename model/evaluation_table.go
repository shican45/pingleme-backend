//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Evaluation_Table 评审表模型
type EvaluationTable struct {
	gorm.Model
	HomeworkID		int	`gorm:"type:int;not null"`
	TeamID			int	`gorm:"type:int;not null"`
}

// Evaluation_Item 评审表项模型
type EvaluationItem struct {
	gorm.Model
	EvaluationTableID			int		`gorm:"type:int;not null"`
	Content						string	`gorm:"type:varchar(255);not null"`
	Score						int		`gorm:"type:int;not null"`
	Suggest						string	`gorm:"type:text"`
	ParentItemID				int		`gorm:"type:int;not null"`
	Sequence					int		`gorm:"type:int;not null"`
}