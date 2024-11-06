package taskrepos

import (
	"context"
	"wark/modules/tasks/taskmodels"

	"github.com/jmoiron/sqlx"
)

type updateTaskRepo struct {
	db *sqlx.DB
}

func NewUpdateTaskRepo(db *sqlx.DB) *updateTaskRepo { return &updateTaskRepo{db} }

func (repo *updateTaskRepo) UpdateTask(task *taskmodels.Task) (int64, error) {
	ctx := context.Background()
	tx, err := repo.db.BeginTxx(ctx, nil)

	if err != nil {
		tx.Rollback()
		panic(err)
	}

	stmt, err := tx.PrepareNamed(`
		UPDATE tasks
			SET (title, description,
				due_date, priority_id,
				task_status, updated_at)
			= (
				:title, :description,
				:due_date, :priority_id,
				:task_status, CURRENT_TIMESTAMP
			)
		WHERE id = :id
	`)

	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(task)

	if err != nil {
		tx.Rollback()
		panic(err)
	}

	r, err := result.RowsAffected()

	if err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()

	return r, nil
}
