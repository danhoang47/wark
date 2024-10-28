package priorityrepos

import (
	"context"
	"database/sql"
	"time"
	"wark/modules/priorities/prioritymodels"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type addPriorityRepo struct {
	db *sqlx.DB
}

func NewAddPriorityRepo(db *sqlx.DB) *addPriorityRepo { return &addPriorityRepo{db} }

func (repo *addPriorityRepo) AddPriority(userId string, priority *prioritymodels.CreatePriority) (int64, error) {
	txCtx, txCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer txCancel()

	tx, err := repo.db.BeginTxx(txCtx, &sql.TxOptions{})

	if err != nil {
		panic(err)
	}

	uuid, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	stmtCtx, stmtCancel := context.WithTimeout(txCtx, time.Second*2)
	defer stmtCancel()

	stmt, err := tx.PrepareNamedContext(stmtCtx, `
		INSERT INTO priorities 
			VALUES(:id, :creator_id, :point, :title)
	`)

	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(map[string]interface{}{
		"id":         uuid,
		"creator_id": userId,
		"point":      priority.Point,
		"title":      priority.Title,
	})

	if err != nil {
		panic(err)
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return result.RowsAffected()
}
