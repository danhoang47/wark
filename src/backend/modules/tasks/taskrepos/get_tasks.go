package taskrepos

import (
	"wark/modules/tasks/taskmodels"

	"github.com/jmoiron/sqlx"
)

type GetTasksRepository interface {
	GetTasks(taskmodels.GetTaskConds) []taskmodels.Task
}

type getTasksRepo struct {
	db *sqlx.DB
}

func NewGetTasksRepo(db *sqlx.DB) *getTasksRepo { return &getTasksRepo{db} }

func (repo *getTasksRepo) GetTasks(conds *taskmodels.GetTaskConds) []taskmodels.Task {
	tasks := []taskmodels.Task{}

	err := repo.db.Select(tasks, `
		SELECT * 
			FROM tasks
			WHERE id < $1
			LIMIT 20
	`, conds.NextCursor)

	if err != nil {
		panic(err)
	}

	return tasks
}
