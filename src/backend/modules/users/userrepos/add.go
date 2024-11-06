package userrepos

import (
	"context"
	"database/sql"
	"errors"
	"wark/common"
	"wark/modules/users/usermodels"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jmoiron/sqlx"
)

type addUserRepo struct {
	db *sqlx.DB
}

func NewAddUserRepo(db *sqlx.DB) *addUserRepo { return &addUserRepo{db} }

func (repo *addUserRepo) Add(createUser *usermodels.CreateUser) (int, error) {
	txCtx := context.Background()

	tx, err := repo.db.BeginTxx(txCtx, nil)

	if err != nil {
		tx.Rollback()
		panic(err)
	}

	row := tx.QueryRowx(`SELECT 1 as "result" FROM users WHERE username = $1`, createUser.Username)
	var result int

	if err := row.Err(); err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := row.Scan(&result); !errors.Is(err, sql.ErrNoRows) || result == 1 {
		tx.Rollback()
		return 0, err
	}

	id, _ := uuid.NewV7()
	user := usermodels.User{
		SQLModel: common.SQLModel{
			Id:     id,
			Status: pgtype.Bits{Bytes: []byte{1}, Len: 1, Valid: true},
		},
		Username: createUser.Username,
		Password: createUser.Password,
		Salt:     createUser.Salt,
	}
	addUserResult, err := tx.NamedExec(`
		INSERT 
			INTO users(id, username, salt, password, status)
			VALUES(:id, :username, :salt, :password, :status)
	`, user)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	rowsAffected, err := addUserResult.RowsAffected()

	if err != nil || rowsAffected == 0 {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return int(rowsAffected), nil
}
