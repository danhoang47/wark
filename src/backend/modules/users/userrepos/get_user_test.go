package userrepos

import (
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestGetUserWithConds(t *testing.T) {
	db, err := sqlx.Open("pgx", "postgres://postgres:0000@localhost:5432/wark")

	if err != nil {
		t.Fatal(err)
	}

	getOneUserRepo := NewGetUserRepo(db)
	user, err := getOneUserRepo.GetUser(&GetOneUserOptions{
		Id: "01928b50-710b-74a1-8723-4f762c0dd1d0",
	})

	if err != nil {
		t.Fatal(err)
	}

	if user == nil {
		t.Fatalf("TestGetUserWithConds: expect 1, have 0")
	}
}
