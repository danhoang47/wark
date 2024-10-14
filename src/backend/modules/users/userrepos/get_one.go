package userrepos

import (
	"database/sql"
	"errors"
	"strings"
	"wark/modules/users/usermodels"

	"github.com/jmoiron/sqlx"
)

type getOneUserRepo struct {
	db *sqlx.DB
}

func NewGetOneUserRepo(db *sqlx.DB) *getOneUserRepo { return &getOneUserRepo{db} }

type GetOneUserOptions struct {
	Id       string
	Username string
	Status   int
}

func (repo *getOneUserRepo) GetOne(options *GetOneUserOptions) (*usermodels.User, error) {
	query := strings.Builder{}
	conds := strings.Builder{}

	query.WriteString(`SELECT * FROM users`)

	if options.Id != "" {
		conds.WriteString(`id=:id`)
	}

	if options.Username != "" {
		conds.WriteString(`username=:username`)
	}

	if options.Status != 0 {
		conds.WriteString(`status=bit(:status)`)
	}

	if conds.Len() != 0 {
		strConds := conds.String()
		strConds = strings.ReplaceAll(strConds, ";", " AND ")
		query.WriteString("WHERE " + strConds)
	}

	stmt, err := repo.db.PrepareNamed(query.String())

	if err != nil {
		return nil, err
	}

	var user usermodels.User

	row := stmt.QueryRowx(options)

	if err := row.Err(); err != nil {
		return nil, err
	}

	if err := row.StructScan(&user); !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &user, nil
}
