package categoryrepos

import (
	"testing"
	"wark/modules/categories/categorymodels"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestAddCategory(t *testing.T) {
	db, err := sqlx.Open("pgx", "postgres://postgres:0000@localhost:5432/wark")

	if err != nil {
		t.Fatal(err)
	}

	repo := NewAddCategoryRepo(db)

	category := categorymodels.CreateCategory{
		Title: "Critical",
		Color: "#FFF",
		Icon:  "faSmile",
	}

	rowsAffected, err := repo.AddCategory("0192909b-f6bd-747f-8db7-4690486be5b2", &category)

	if err != nil {
		t.Fatal("failed with error ", err)
	}

	if rowsAffected != 1 {
		t.Fatal("expect rowsAffected to be 1, got ", rowsAffected)
	}
}
