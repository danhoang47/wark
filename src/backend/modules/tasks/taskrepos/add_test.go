package taskrepos

import (
	"testing"
	"time"
	"wark/modules/tasks/taskmodels"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestAddTask(t *testing.T) {
	db, err := sqlx.Open("pgx", "postgres://postgres:0000@localhost:5432/wark")

	if err != nil {
		t.Fatal(err)
	}

	addTaskRepo := NewAddTaskRepo(db)

	uuid, err := uuid.NewV7()

	if err != nil {
		t.Fatal(err)
	}

	task := &taskmodels.CreateTask{
		CreatorId:   uuid,
		ParentId:    uuid,
		Title:       "Test_task_2",
		Description: "Deserunt laborum do irure consectetur.",
		DueDate:     time.Now(),
		PriorityId:  uuid,
	}

	r, err := addTaskRepo.AddTask(task)

	if err != nil {
		t.Fatal(err)
	}

	if r != 1 {
		t.Fatalf("expect result is 1, got %v", r)
	}
}
