//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TestRepository struct {
	repo *Repository
	mock sqlmock.Sqlmock
	db   *sql.DB
}

func (tRepo *TestRepository) InitTest() error {
	db, mock, err := sqlmock.New()
	if err != nil {
		return err
	}
	tRepo.mock = mock

	r := new(Repository)
	r.DB, err = gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	if err != nil {
		return err
	}
	tRepo.repo = r
	tRepo.db = db
	return nil
}
