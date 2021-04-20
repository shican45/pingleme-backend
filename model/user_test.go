//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"time"
)

func TestRepository_GetUser(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	t.Run("GetUser", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "user_name", "password_digest", "role"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), "test1", "password", "role1")

		tRepo.mock.ExpectQuery("SELECT *").WillReturnRows(rows)

		user, err := tRepo.repo.GetUser(1)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, user.UserName, "test1")
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}