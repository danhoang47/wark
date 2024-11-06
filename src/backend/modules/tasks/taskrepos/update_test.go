package taskrepos

import (
	"testing"
	"time"
	"wark/common"
	"wark/modules/tasks/taskmodels"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestUpdateTask(t *testing.T) {
	db, err := sqlx.Open("pgx", "postgres://postgres:0000@localhost:5432/wark")

	if err != nil {
		t.Fatal(err)
	}

	updateTaskRepo := NewUpdateTaskRepo(db)

	id := uuid.MustParse("0192cc3c-2271-73fd-8fe8-c6ece2018e45")
	creatorId := uuid.MustParse("0192909b-f6bd-747f-8db7-4690486be5b2")
	priorityId := uuid.MustParse("0192cc39-d486-7dbd-b968-368698f78126")

	task := &taskmodels.Task{
		SQLModel: common.SQLModel{
			Id: id,
		},
		CreatorId:   creatorId,
		Title:       "Test task 6",
		Description: "Deserunt laborum do irure consectetur.",
		DueDate:     time.Now(),
		PriorityId:  priorityId,
	}

	r, err := updateTaskRepo.UpdateTask(task)

	if err != nil {
		t.Fatal(err)
	}

	if r != 1 {
		t.Fatalf("expect result is 1, got %v", r)
	}
}
