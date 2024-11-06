package taskrepos

import (
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestUpdateTaskCategories(t *testing.T) {
	db, err := sqlx.Open("pgx", "postgres://postgres:0000@localhost:5432/wark")

	if err != nil {
		t.Fatal(err)
	}

	categoryIds := []string{"0192d868-ea5e-7458-ab94-b0beaf7ef508", "0192d874-9098-7cae-a75a-f3f272fbeeed"}
	taskId := "0192cc3c-2271-73fd-8fe8-c6ece2018e45"

	repo := NewUpdateTaskCategories(db)

	if err := repo.UpdateTaskCategories(taskId, categoryIds); err != nil {
		t.Fatal("error occur")
	}
}
