package prioritymodels

import (
	"wark/common"

	"github.com/google/uuid"
)

type Priority struct {
	common.SQLModel
	CreatorId uuid.UUID `json:"creatorId" db:"creator_id"`
	Title     string    `json:"title" db:"title"`
	Point     int8      `json:"point" db:"point"`
}

type CreatePriority struct {
	Title string `json:"title" db:"title"`
	Point int8   `json:"point" db:"point"`
}

type TaskPriority struct {
	Id    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
	Point int8      `json:"point" db:"point"`
}
