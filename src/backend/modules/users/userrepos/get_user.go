package userrepos

import (
	"strings"
	"wark/modules/users/usermodels"

	"github.com/jmoiron/sqlx"
)

type getUserRepo struct {
	db *sqlx.DB
}

func NewGetUserRepo(db *sqlx.DB) *getUserRepo { return &getUserRepo{db} }

type GetOneUserOptions struct {
	Id       string
	Username string
	Status   int
}

func (repo *getUserRepo) GetUser(options *GetOneUserOptions) (*usermodels.User, error) {
	query := &strings.Builder{}
	conds := make([]string, 0)

	query.WriteString(`SELECT * FROM users`)

	if options.Id != "" {
		conds = append(conds, `id=:id`)
	}

	if options.Username != "" {
		conds = append(conds, `username=:username`)
	}

	if options.Status != 0 {
		conds = append(conds, `status=bit(:status)`)
	}

	if len(conds) != 0 {
		strConds := strings.Join(conds, " AND ")
		query.WriteString(" WHERE " + strConds)
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

	if err := row.StructScan(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
