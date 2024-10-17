package taskbiz

import "wark/modules/tasks/taskmodels"

type AddTaskBusiness interface {
	AddTask(*taskmodels.CreateTask) bool
}

type AddTaskRepo interface {
	AddTask(*taskmodels.CreateTask) (int64, error)
}

type addTaskBiz struct {
	taskRepo AddTaskRepo
}

func NewAddTaskRepo(taskRepo AddTaskRepo) AddTaskBusiness { return &addTaskBiz{taskRepo} }

func (biz *addTaskBiz) AddTask(task *taskmodels.CreateTask) bool {
	r, err := biz.taskRepo.AddTask(task)

	if err != nil {
		panic(err)
	}

	return r == 1
}
