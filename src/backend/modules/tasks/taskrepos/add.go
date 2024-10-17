package taskrepos

import (
	"context"
	"log"
	"wark/modules/tasks/taskmodels"

	"github.com/jmoiron/sqlx"
)

type addTaskRepo struct {
	db *sqlx.DB
}

func NewAddTaskRepo(db *sqlx.DB) *addTaskRepo { return &addTaskRepo{db} }

func (repo *addTaskRepo) AddTask(task *taskmodels.CreateTask) (int64, error) {
	ctx := context.Background()
	tx, err := repo.db.BeginTxx(ctx, nil)

	if err != nil {
		tx.Rollback()
		panic(err)
	}

	stmt, err := tx.PrepareNamed(`
		INSERT 
			INTO tasks(id, creator_id, parent_id,
				title, description,
				due_date, priority_id,
				task_status, status)
			VALUES(
				:id, :creator_id, :parent_id,
				:title, :description,
				:due_date, :priority_id,
				:task_status, :status
			)
	`)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	t := task.ToTask(1)

	result, err := stmt.Exec(t)

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
