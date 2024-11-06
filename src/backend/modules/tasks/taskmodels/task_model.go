package taskmodels

import (
	"time"
	"wark/common"
	"wark/modules/categories/categorymodels"
	"wark/modules/priorities/prioritymodels"

	"github.com/google/uuid"
)

type Task struct {
	common.SQLModel
	CreatorId   uuid.UUID     `json:"creatorId" db:"creator_id"`
	ParentId    uuid.NullUUID `json:"parentId" db:"parent_id"`
	Title       string        `json:"title" db:"title"`
	Description string        `json:"description" db:"description"`
	DueDate     time.Time     `json:"dueDate" db:"due_date"`
	PriorityId  uuid.UUID     `json:"priorityId" db:"priority_id"`
	TaskStatus  int8          `json:"taskStatus" db:"task_status"`
}

type CreateTask struct {
	CreatorId   uuid.UUID     `json:"creatorId" db:"creator_id"`
	ParentId    uuid.NullUUID `json:"parentId" db:"parent_id"`
	Title       string        `json:"title" db:"title"`
	Description string        `json:"description" db:"description"`
	DueDate     time.Time     `json:"dueDate" db:"due_date"`
	PriorityId  uuid.UUID     `json:"priorityId" db:"priority_id"`
	TaskStatus  int8          `json:"taskStatus" db:"task_status"`
}

func (task *CreateTask) ToTask(status byte) *Task {
	uuid, err := uuid.NewV7()

	if err != nil {
		panic(err)
	}

	t := &Task{
		SQLModel: common.SQLModel{
			Id: uuid,
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

type AggregatedTask struct {
	common.SQLModel
	ParentId    uuid.NullUUID                `json:"parentId,omitempty" db:"parent_id"`
	Title       string                       `json:"title" db:"title"`
	Description string                       `json:"description" db:"description"`
	DueDate     time.Time                    `json:"dueDate" db:"due_date"`
	TaskStatus  int8                         `json:"taskStatus" db:"task_status"`
	Priority    prioritymodels.TaskPriority  `json:"priority"`
	Categories  []categorymodels.GetCategory `json:"categories"`
}
