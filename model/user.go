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
	Roles          []Role `gorm:"many2many:user_role"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

type UserRepositoryInterface interface {
	GetUser(ID interface{}) (User, error)
	GetUserByUID(UID string) (User, error)
}

// GetUser 用ID获取用户
func (Repo *Repository) GetUser(ID interface{}) (User, error) {
	var user User
	result := Repo.DB.First(&user, ID)
	return user, result.Error
}

// GetUserByUID 用UID获取用户
func (Repo *Repository) GetUserByUID(UID string) (User, error) {
	var user User
	result := Repo.DB.Where("uid = ?", UID).First(&user)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
