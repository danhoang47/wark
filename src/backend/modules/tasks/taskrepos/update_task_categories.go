package taskrepos

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type updateTasksCategories struct {
	db *sqlx.DB
}

func NewUpdateTaskCategories(db *sqlx.DB) *updateTasksCategories {
	return &updateTasksCategories{db}
}

func (repo *updateTasksCategories) UpdateTaskCategories(taskId string, categoryIds []string) error {
	txCtx, txCtxCancel := context.WithTimeout(context.Background(), time.Second*3)
	defer txCtxCancel()

	tx, err := repo.db.BeginTxx(txCtx, &sql.TxOptions{})

	if err != nil {
		panic(err)
	}

	var currentCategoryIds []string = []string{}

	err = tx.Select(
		&currentCategoryIds,
		`SELECT category_id 
			FROM tasks_categories 
			WHERE task_id = $1`,
		taskId,
	)

	if err != nil {
		panic(err)
	}

	var deletedCategoryIds []string

	for _, categoryId := range currentCategoryIds {
		if !slices.Contains(categoryIds, categoryId) {
			deletedCategoryIds = append(deletedCategoryIds, categoryId)
		}
	}

	var addedCategoryIds []string

	for _, categoryId := range categoryIds {
		if !slices.Contains(currentCategoryIds, categoryId) {
			addedCategoryIds = append(addedCategoryIds, categoryId)
		}
	}

	if len(addedCategoryIds) > 0 {

		sql := strings.Builder{}
		values := make([]string, 0, len(addedCategoryIds))

		sql.WriteString(`INSERT INTO tasks_categories VALUES `)

		for _, categoryId := range addedCategoryIds {
			values = append(values, fmt.Sprintf("('%s', '%s')", taskId, categoryId))
		}

		sql.WriteString(strings.Join(values, ","))

		log.Println(sql.String())

		insertStmt, err := tx.Preparex(sql.String())

		if err != nil {
			panic(err)
		}

		stmtCtx, stmtCtxCancel := context.WithTimeout(txCtx, time.Second*2)
		defer stmtCtxCancel()

		_, err = insertStmt.ExecContext(stmtCtx)

		if err != nil {
			panic(err)
		}
	}

	if len(deletedCategoryIds) > 0 {
		deleteInSql, args, err := sqlx.In(`
		DELETE FROM tasks_categories 
		WHERE task_id = $1
		AND category_id IN ($2) 
	`, taskId, deletedCategoryIds)

		if err != nil {
			panic(err)
		}

		deleteInSql = repo.db.Rebind(deleteInSql)
		_, err = repo.db.Exec(deleteInSql, args...)

		if err != nil {
			panic(err)
		}
	}

	tx.Commit()

	return nil
}
