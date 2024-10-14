package userrepos

import (
	"errors"
	"testing"
	"wark/common"
	"wark/components/hasher"
	"wark/modules/users/usermodels"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestCreateUser_WithFullData(t *testing.T) {
	db, err := sqlx.Open("pgx", "postgres://postgres:0000@localhost:5432/wark")

	if err != nil {
		t.Fatal(err)
	}

	salt := common.GetSalt()
	hasher := hasher.New()
	hashedPassword := hasher.Hash([]byte("0000" + salt))

	testCreateUser := &usermodels.CreateUser{
		Username: "dathnq",
		Salt:     salt,
		Password: hashedPassword,
	}

	addUserRepo := NewAddUserRepo(db)
	rowsAffected, err := addUserRepo.Add(testCreateUser)

	if err != nil || rowsAffected != 1 {
		t.Fatal(err)
	}
}

func TestCreateUser_DuplicateUsername(t *testing.T) {
	db, err := sqlx.Open("pgx", "postgres://postgres:0000@localhost:5432/wark")

	if err != nil {
		t.Fatal(err)
	}

	salt := common.GetSalt()
	hasher := hasher.New()
	hashedPassword := hasher.Hash([]byte("0000" + salt))

	testCreateUser := &usermodels.CreateUser{
		Username: "dathnq",
		Salt:     salt,
		Password: hashedPassword,
	}

	addUserRepo := NewAddUserRepo(db)
	_, err = addUserRepo.Add(testCreateUser)

	if err != nil && !errors.Is(err, common.ErrUsernameHasTaken) {
		t.Fatal(err)
	}
}
