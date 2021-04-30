package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHomework(t *testing.T) {
	var tRepo TestRepository
	err := tRepo.InitTest()
	if err != nil {
		t.Error(err)
	}
	defer tRepo.db.Close()

	t.Run("GetHomeworkByID", func(t *testing.T) {
		tRepo.mock.ExpectQuery("SELECT (.+) FROM `homework` WHERE `homework`.`id` = \\? AND (.*)").
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "class_id", "type",
				"title", "content", "start_time", "end_time"}).
				AddRow(1, time.Now(), time.Now(), time.Now(), 1, 1, "title_test", "content_test", time.Now(), time.Now()))
		homework, err := tRepo.repo.GetHomeworkByID(1)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "title_test", homework.Title)
		assert.Equal(t, "content_test", homework.Content)

		if err := tRepo.mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}
