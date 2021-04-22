//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
	"time"
)

// Homework 作业模型
type Homework struct {
	gorm.Model
	ClassID   int       `gorm:"type:int;not null"`
	Type      uint8     `gorm:"type:int;not null"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Content   string    `gorm:"type:text;not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
}

// Scoring_Item 评分项模型
type ScoringItem struct {
	gorm.Model
	HomeworkID   int    `gorm:"type:int;not null"`
	Description  string `gorm:"type:varchar(255);not null"`
	Score        int    `gorm:"type:int;not null"`
	Option       uint8  `gorm:"type:int;not null"`
	Note         string `gorm:"type:varchar(255)"`
	AssistantID  int    `gorm:"type:int;not null"`
	ParentItemID int    `gorm:"type:int;not null"`
	Sequence     int    `gorm:"type:int;not null"`
}
