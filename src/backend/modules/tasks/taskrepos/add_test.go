package taskrepos

import (
	"strconv"
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

	creatorId, err := uuid.Parse("0192909b-f6bd-747f-8db7-4690486be5b2")

	if err != nil {
		t.Fatal(err)
	}

	priorityId, err := uuid.Parse("0192cc39-d486-7dbd-b968-368698f78126")

	if err != nil {
		t.Fatal(err)
	}

	for index := range 30 {
		task := &taskmodels.CreateTask{
			CreatorId:   creatorId,
			Title:       "Test task " + strconv.Itoa(index),
			Description: "Deserunt laborum do irure consectetur.",
			DueDate:     time.Now(),
			PriorityId:  priorityId,
		}

		_, err := addTaskRepo.AddTask(task)

		if err != nil {
			t.Fatal(err)
		}
	}

	if err != nil {
		t.Fatal(err)
	}
}
