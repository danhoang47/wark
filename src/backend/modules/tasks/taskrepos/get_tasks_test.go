package taskrepos

import (
	"testing"
	"wark/modules/tasks/taskmodels"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestGenerateSqlQuery(t *testing.T) {
	db, err := sqlx.Open("pgx", "postgres://postgres:0000@localhost:5432/wark")

	if err != nil {
		t.Fatal(err)
	}

	getTasksRepo := NewGetTasksRepo(db)

	if err != nil {
		panic(err)
	}

	tasks := getTasksRepo.GetTasks("0192909b-f6bd-747f-8db7-4690486be5b2", &taskmodels.GetTaskConds{})

	if len(tasks) == 0 {
		t.Fatal("expect tasks not empty")
	}
}
