//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
)

// Role 角色模型
type Role struct {
	gorm.Model
	Type				uint8	`gorm:"type:int;not null;unique"`
	Permission 			[]Permission	`gorm:"many2many:role_permission;"`
}

// Permission 权限模型
type Permission struct {
	gorm.Model
	Type				uint8	`gorm:"type:int;not null;unique"`
	Role 				[]Role	`gorm:"many2many:role_permission;"`
}