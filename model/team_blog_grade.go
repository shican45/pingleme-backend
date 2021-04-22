//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Team_Blog_Grade 团队博客成绩模型
type TeamBlogGrade struct {
	gorm.Model
	ScoringItemID		int	`gorm:"type:int;not null"`
	ScorekeeperID		int	`gorm:"type:int;not null"`
	Grade				int	`gorm:"type:int;not null"`
}