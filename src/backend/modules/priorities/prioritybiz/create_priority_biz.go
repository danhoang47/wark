package prioritybiz

import "wark/modules/priorities/prioritymodels"

type CreatePriorityRepo interface {
	AddPriority(string, *prioritymodels.CreatePriority) (int64, error)
}

type CreatePriorityBiz interface {
	CreatePriority(string, *prioritymodels.CreatePriority) (int64, error)
}

type createPriorityBiz struct {
	repo CreatePriorityRepo
}

func NewCreatePriorityBiz(repo CreatePriorityRepo) CreatePriorityBiz {
	return &createPriorityBiz{repo}
}

func (biz *createPriorityBiz) CreatePriority(
	userId string,
	priority *prioritymodels.CreatePriority,
) (int64, error) {
	return biz.repo.AddPriority(userId, priority)
}
