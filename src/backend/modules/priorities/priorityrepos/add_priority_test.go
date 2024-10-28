package priorityrepos

import (
	"testing"
	"wark/modules/priorities/prioritymodels"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestAddPriority(t *testing.T) {
	db, err := sqlx.Open("pgx", "postgres://postgres:0000@localhost:5432/wark")

	if err != nil {
		t.Fatal(err)
	}

	repo := NewAddPriorityRepo(db)

	priority := prioritymodels.CreatePriority{
		Title: "Critical",
		Point: 10,
	}

	rowsAffected, err := repo.AddPriority("0192909b-f6bd-747f-8db7-4690486be5b2", &priority)

	if err != nil {
		t.Fatal("failed with error ", err)
	}

	if rowsAffected != 1 {
		t.Fatal("expect rowsAffected to be 1, got ", rowsAffected)
	}
}
