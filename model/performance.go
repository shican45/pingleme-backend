//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Performance 绩效模型
type Performance struct {
	gorm.Model
	HomeworkID				int	`gorm:"type:int;not null"`
	StudentID				int	`gorm:"type:int;not null"`
	Percentage				int	`gorm:"type:int;not null"`
}