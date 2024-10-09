package common

import (
	"time"

	"github.com/google/uuid"
)

const (
	Available = iota
	Deleted
)

type SQLModel struct {
	Id        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	Status    int       `json:"status" db:"status"`
}
