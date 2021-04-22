//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Personal_Blog_Grade 个人博客成绩模型
type PersonalBlogGrade struct {
	gorm.Model
	ScoringItemID		int	`gorm:"type:int;not null"`
	ScorekeeperID		int	`gorm:"type:int;not null"`
	Grade				int	`gorm:"type:int;not null"`
}