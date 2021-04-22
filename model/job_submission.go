//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Job_Submission 作业提交模型
type JobSubmission struct {
	gorm.Model
	SubmitterID  int    `gorm:"type:int;not null"`
	HomeworkID   int    `gorm:"type:int;not null"`
	SubmitStatus uint8  `gorm:"type:int;not null"`
	Filepath     string `gorm:"type:varchar(255)"`
}
