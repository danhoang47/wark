package taskrepos

import (
	"log"
	"strings"
	"wark/modules/categories/categorymodels"
	"wark/modules/tasks/taskmodels"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type getTasksRepo struct {
	db *sqlx.DB
}

func NewGetTasksRepo(db *sqlx.DB) *getTasksRepo { return &getTasksRepo{db} }

func (repo *getTasksRepo) GetTasks(userId string, conds *taskmodels.GetTaskConds) []taskmodels.AggregatedTask {
	var tasks []taskmodels.AggregatedTask

	sqlBuilder := &strings.Builder{}

	sqlBuilder.WriteString(`
		SELECT 
			t.id, t.parent_id, t.title,
			t.description, t.due_date, t.task_status,
			t.created_at, t.updated_at, 
			p.id, p.point, p.title 
		FROM tasks t JOIN priorities p 
			ON p.id = t.priority_id
		WHERE t.creator_id = :userId
	`)

	if conds.NextCursor.Valid {
		sqlBuilder.WriteString(` AND t.id < :nextCursor`)
	} else {
		log.Println("nextCursor not provided")
	}

	sqlBuilder.WriteString(` ORDER BY t.updated_at DESC LIMIT 20`)

	query := sqlBuilder.String()

	log.Println(query)

	namedStmt, err := repo.db.PrepareNamed(query)

	if err != nil {
		panic(err)
	}

	rows, err := namedStmt.Queryx(map[string]interface{}{
		"userId":     userId,
		"nextCursor": conds.NextCursor.UUID.String(),
	})

	if err != nil {
		log.Println(err)
		panic(err)
	}

	for rows.Next() {
		task := taskmodels.AggregatedTask{}

		if err := rows.Scan(
			&task.Id, &task.ParentId, &task.Title,
			&task.Description, &task.DueDate,
			&task.TaskStatus, &task.CreatedAt, &task.UpdatedAt,
			&task.Priority.Id, &task.Priority.Point, &task.Priority.Title,
		); err != nil {
			panic(err)
		}

		tasks = append(tasks, task)
	}

	log.Println("task len: ", len(tasks))

	if len(tasks) == 0 {
		return []taskmodels.AggregatedTask{}
	}

	ids := make([]uuid.UUID, len(tasks))

	for _, task := range tasks {
		ids = append(ids, task.Id)
	}

	query, args, err := sqlx.In(`
		SELECT t.task_id, c.id, c.title, c.color, c.icon  FROM tasks_categories t
		JOIN categories c ON t.category_id = c.id
		WHERE task_id IN (?)
	`, ids)

	if err != nil {
		panic(err)
	}

	query = repo.db.Rebind(query)
	rows, err = repo.db.Queryx(query, args...)

	if err != nil {
		rows.Close()
		panic(err)
	}

	for rows.Next() {
		category := categorymodels.GetCategory{}
		var taskId uuid.UUID

		if err := rows.Scan(&taskId, &category.Id, &category.Title, &category.Color, &category.Icon); err != nil {
			rows.Close()
			panic(err)
		}

		for _, task := range tasks {
			if strings.Compare(task.Id.String(), taskId.String()) == 0 {
				task.Categories = append(task.Categories, category)
			}
		}
	}

	return tasks
}
