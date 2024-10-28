package taskbiz

import "wark/modules/tasks/taskmodels"

type GetTasksRepo interface {
	GetTasks(string, *taskmodels.GetTaskConds) []taskmodels.AggregatedTask
}

type GetTasksBiz interface {
	GetTasks(string, *taskmodels.GetTaskConds) []taskmodels.AggregatedTask
}

type getTasksBiz struct {
	repo GetTasksRepo
}

func NewGetTasksBiz(repo GetTasksRepo) GetTasksBiz { return &getTasksBiz{repo} }

func (biz *getTasksBiz) GetTasks(userId string, conds *taskmodels.GetTaskConds) []taskmodels.AggregatedTask {
	return biz.repo.GetTasks(userId, conds)
}
