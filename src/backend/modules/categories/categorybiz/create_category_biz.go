package categorybiz

import "wark/modules/categories/categorymodels"

type CreateCategoryRepo interface {
	AddCategory(string, *categorymodels.CreateCategory) (int64, error)
}

type CreateCategoryBiz interface {
	CreateCategory(string, *categorymodels.CreateCategory) (int64, error)
}

type createCategoryBiz struct {
	repo CreateCategoryRepo
}

func NewCreateCategoryBiz(repo CreateCategoryRepo) CreateCategoryBiz {
	return &createCategoryBiz{repo}
}

func (biz *createCategoryBiz) CreateCategory(
	userId string,
	category *categorymodels.CreateCategory,
) (int64, error) {
	return biz.repo.AddCategory(userId, category)
}
