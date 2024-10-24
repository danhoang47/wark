package taskmodels

import (
	"time"
	"wark/common"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Task struct {
	common.SQLModel
	CreatorId   uuid.UUID `json:"creatorId" db:"creator_id"`
	ParentId    uuid.UUID `json:"parentId" db:"parent_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	DueDate     time.Time `json:"dueDate" db:"due_date"`
	PriorityId  uuid.UUID `json:"priorityId" db:"priority_id"`
	TaskStatus  int8      `json:"taskStatus" db:"task_status"`
}

type CreateTask struct {
	CreatorId   uuid.UUID `json:"creatorId" db:"creator_id"`
	ParentId    uuid.UUID `json:"parentId" db:"parent_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	DueDate     time.Time `json:"dueDate" db:"due_date"`
	PriorityId  uuid.UUID `json:"priorityId" db:"priority_id"`
	TaskStatus  int8      `json:"taskStatus" db:"task_status"`
}

func (task *CreateTask) ToTask(status byte) *Task {
	uuid, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	t := &Task{
		SQLModel: common.SQLModel{
			Id: uuid,
			Status: &pgtype.Bits{
				Bytes: []byte{status},
				Len:   1,
				Valid: true,
			},
		},
		CreatorId:   task.CreatorId,
		ParentId:    task.ParentId,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		PriorityId:  task.PriorityId,
		TaskStatus:  task.TaskStatus,
	}

	return t
}

type GetTaskConds struct {
	common.Paging
}
