//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Partner 结对模型
type Partner struct {
	gorm.Model
	Student1ID int `gorm:"type:int;not null"`
	Student2ID int `gorm:type:int`
}
