package userbiz

import (
	"testing"
	"wark/components/hasher"
	"wark/modules/users/usermodels"
	"wark/modules/users/userrepos"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func TestUserBiz_LoginWithValidCredential(t *testing.T) {
	u := &usermodels.CreateUser{
		Username: "dathnq1",
		Password: "123456",
	}

	db, err := sqlx.Open("pgx", "postgres://postgres:0000@localhost:5432/wark")

	if err != nil {
		t.Fatal(err)
	}

	getUserRepo := userrepos.NewGetUserRepo(db)
	biz := NewLoginUserBiz(getUserRepo)

	user, err := biz.Login(hasher.New(), u)

	if err != nil {
		t.Fatal(err)
	}

	if user == nil {
		t.Fatalf("expect user with %v, but got none", u)
	}
}
