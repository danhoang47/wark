package common

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const (
	Available = true
	Deleted   = false
)

type SQLModel struct {
	Id        uuid.UUID   `json:"id" db:"id"`
	CreatedAt time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time   `json:"updatedAt" db:"updated_at"`
	Status    pgtype.Bits `json:"-" db:"status"`
}

type Paging struct {
	Cursor     uuid.UUID     `json:"cursor,omitempty"`
	NextCursor uuid.NullUUID `json:"nextCursor,omitempty"`
}
