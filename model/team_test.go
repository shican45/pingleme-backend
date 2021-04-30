//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"time"
)

func TestTeam(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	t.Run("GetTeam", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "number", "name", "groupLeaderID", "classID"}).
			AddRow(1, time.Now(), time.Now(), time.Now(), 1, "pingleme", 221801219, 1)

		tRepo.mock.ExpectQuery("SELECT *").WillReturnRows(rows)

		team, err := tRepo.repo.GetTeam(1)

		if err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, team.Name, "pingleme")
		}

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})


	//t.Run("SetTeam", func(t *testing.T) {
	//	team := Team{
	//		Number:        1,
	//		Name:          "pingleme",
	//		GroupLeaderID: 1,
	//		ClassID:       1,
	//	}
	//	sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "number", "name", "groupLeaderID", "classID"}).
	//		AddRow(1, time.Now(), time.Now(), time.Now(), 1, "pingleme", 221801219, 1)
	//	tRepo.mock.ExpectBegin()
	//	tRepo.mock.ExpectExec("INSERT INTO team").
	//		WithArgs(1, time.Now(),time.Now(),time.Now(),1,"pingleme",1,1)
	//	//	WillReturnResult(sqlmock.NewResult(1, 1))
	//	//tRepo.mock.ExpectExec("UPDATE Team").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	//	tRepo.mock.ExpectRollback()
	//	rowsAffected, err := tRepo.repo.SetTeam(team)
	//
	//	if err != nil {
	//		t.Error(err)
	//	} else {
	//		assert.Equal(t, rowsAffected, 1)
	//	}
	//
	//	if err := tRepo.mock.ExpectationsWereMet(); err != nil {
	//		t.Error(err)
	//	}
	//})
	//t.Run("SetClassNameByID", func(t *testing.T) {
	//	sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "number", "name", "groupLeaderID", "classID"}).
	//		AddRow(1, time.Now(), time.Now(), time.Now(), 1, "pingleme", 221801219, 1)
	//
	//	tRepo.mock.ExpectExec("UPDATE Team").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	//	rowsAffected, err := tRepo.repo.SetClassNameByID(1,"pinggepi")
	//
	//	if err != nil {
	//		t.Error(err)
	//	} else {
	//		assert.Equal(t, rowsAffected, 1)
	//	}
	//
	//	if err := tRepo.mock.ExpectationsWereMet(); err != nil {
	//		t.Error(err)
	//	}
	//})


}
