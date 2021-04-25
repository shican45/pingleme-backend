//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// PersonalBlogScore 个人博客成绩模型
type PersonalBlogScore struct {
	gorm.Model
	ScoringItemID int `gorm:"type:int;not null"`
	ScorekeeperID int `gorm:"type:int;not null"`
	Grade         int `gorm:"type:int;not null"`
}

// TeamBlogScore 团队博客成绩模型
type TeamBlogScore struct {
	gorm.Model
	ScoringItemID int `gorm:"type:int;not null"`
	ScorekeeperID int `gorm:"type:int;not null"`
	Grade         int `gorm:"type:int;not null"`
}
