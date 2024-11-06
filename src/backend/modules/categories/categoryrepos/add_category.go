package categoryrepos

import (
	"context"
	"database/sql"
	"time"
	"wark/modules/categories/categorymodels"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type addCategoryRepo struct {
	db *sqlx.DB
}

func NewAddCategoryRepo(db *sqlx.DB) *addCategoryRepo { return &addCategoryRepo{db} }

func (repo *addCategoryRepo) AddCategory(userId string, category *categorymodels.CreateCategory) (int64, error) {
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
		INSERT INTO categories 
			VALUES(:id, :creator_id, :title, :color, :icon)
	`)

	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(map[string]interface{}{
		"id":         uuid,
		"creator_id": userId,
		"title":      category.Title,
		"color":      category.Color,
		"icon":       category.Icon,
	})

	if err != nil {
		panic(err)
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return result.RowsAffected()
}
