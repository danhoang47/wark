package jobmodels

import (
	"wark/common"

	"github.com/google/uuid"
)

const (
	Once = iota
	Interval
)

// boolean ops
const (
	Equal = iota
	LessThan
	GreaterThan
	NotEqual
)

// conjunctions
const (
	And = iota
	Or  = iota
)

// actions
const (
	Sort = iota
	Update
	Delete
)

type JobTime struct {
	Type int    `json:"type"`
	Time string `json:"time"`
}

type JobCondition struct {
	Operator int         `json:"operator"`
	Field    string      `json:"field"`
	Value    interface{} `json:"value"`
}

type JobConditions struct {
	Conditions   []JobCondition `json:"conditions"`
	Conjunctions []int          `json:"conjunctions"`
}

type JobAction struct {
	Action  int    `json:"action"`
	Field   string `json:"field"`
	Order   string `json:"order,omitempty"`
	OrderBy bool   `json:"orderBy,omitempty"`
}

type JobBody struct {
	Time       JobTime       `json:"time"`
	Conditions JobConditions `json:"conditions"`
	Action     JobAction     `json:"action"`
}

type Job struct {
	common.SQLModel
	CreatorId uuid.UUID `json:"creator_id" db:"creator_id"`
	Title     string    `json:"title" db:"title"`
	Body      JobBody   `json:"body" db:"body"`
}

type CreateJob struct {
	CreatorId uuid.UUID `json:"creator_id" db:"creator_id"`
	Title     string    `json:"title" db:"title"`
	Body      JobBody   `json:"body" db:"body"`
}
