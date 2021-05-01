//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"time"
)

func TestPerformance(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	t.Run("GetPerformance", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "homeworkID", "studentID", "percentage"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), 1, 1, 1)

		tRepo.mock.ExpectQuery("SELECT *").WillReturnRows(rows)

		performance, err := tRepo.repo.GetPerformance(1)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, performance.Percentage, 1)
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})

}
