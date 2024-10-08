package common

import "time"

type GUID string

const (
	Available = iota
	Deleted
)

type CommonSQL struct {
	Id        GUID      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Status    int       `json:"status"`
}
