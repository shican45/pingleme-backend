//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Evaluation_Item_Grade 评审表项成绩模型
type EvaluationItemGrade struct {
	gorm.Model
	ScoringItemID		int		`gorm:"type:int;not null"`
	TeamID				int		`gorm:"type:int;not null"`
	UID					string	`gorm:"type:varchar(9);not null"`
	Grade				int		`gorm:"type:int;not null"`
}