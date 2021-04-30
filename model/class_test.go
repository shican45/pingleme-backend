package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClass(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	t.Run("GetClassByID", func(t *testing.T) {
		tRepo.mock.ExpectQuery("SELECT (.+) FROM `class` WHERE `class`.`id` = \\? AND (.*)").
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"}).
				AddRow(1, time.Now(), time.Now(), time.Now(), "software21class"))
		class, err := tRepo.repo.GetClassByID(1)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "software21class", class.Name)

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}
