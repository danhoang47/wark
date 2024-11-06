package categorymodels

import (
	"wark/common"

	"github.com/google/uuid"
)

type Category struct {
	common.SQLModel
	CreatorId uuid.UUID `json:"creatorId" db:"creator_id"`
	Title     string    `json:"title" db:"title"`
	Color     string    `json:"color" db:"color"`
	Icon      string    `json:"icon" db:"icon"`
}

type CreateCategory struct {
	Title string `json:"title" db:"title"`
	Color string `json:"color" db:"color"`
	Icon  string `json:"icon" db:"icon"`
}

type GetCategory struct {
	Id    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
	Color string    `json:"color" db:"color"`
	Icon  string    `json:"icon" db:"icon"`
}
