//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UID            string `gorm:"type:varchar(9);not null;unique"`
	PasswordDigest string `gorm:"type:varchar(16);not null"`
	Nickname       string `gorm:"type:varchar(20);not null;unique"`
	Role           uint8  `gorm:"type:int;not null"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

// GetUser 用ID获取用户
//func GetUser(ID interface{}) (User, error) {
//	var user User
//	result := DB.First(&user, ID)
//	return user, result.Error
//}
func (Repo *Repository) GetUser(ID interface{}) (User, error) {
	var user User
	result := Repo.DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (assistant *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	assistant.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (assistant *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(assistant.PasswordDigest), []byte(password))
	return err == nil
}
