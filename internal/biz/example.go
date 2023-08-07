package biz

import (
	"context"

	v1 "github.com/ZQCard/kbk-layout/api/example/v1"
	"github.com/ZQCard/kbk-layout/internal/domain"
	"github.com/go-kratos/kratos/v2/log"
)

type ExampleRepo interface {
	ListExample(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*domain.Example, error)
	GetExampleCount(ctx context.Context, params map[string]interface{}) (int64, error)
	GetExample(ctx context.Context, params map[string]interface{}) (*domain.Example, error)
	CreateExample(ctx context.Context, example *domain.Example) (*domain.Example, error)
	UpdateExample(ctx context.Context, example *domain.Example) error
	DeleteExample(ctx context.Context, example *domain.Example) error
	RecoverExample(ctx context.Context, example *domain.Example) error
}

type ExampleUsecase struct {
	repo   ExampleRepo
	logger *log.Helper
}

func NewExampleUsecase(repo ExampleRepo, logger log.Logger) *ExampleUsecase {
	return &ExampleUsecase{repo: repo, logger: log.NewHelper(log.With(logger, "module", "usecase/example"))}
}

func (suc *ExampleUsecase) ListExample(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*domain.Example, int64, error) {
	list, err := suc.repo.ListExample(ctx, page, pageSize, params)
	if err != nil {
		return nil, 0, err
	}
	count, err := suc.repo.GetExampleCount(ctx, params)
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (suc *ExampleUsecase) GetExample(ctx context.Context, example *domain.Example) (*domain.Example, error) {
	params := make(map[string]interface{})
	params["id"] = example.Id
	return suc.repo.GetExample(ctx, params)
}

func (suc *ExampleUsecase) CreateExample(ctx context.Context, example *domain.Example) (*domain.Example, error) {
	count, _ := suc.repo.GetExampleCount(ctx, map[string]interface{}{
		"name":       example.Name,
		"id_deleted": true,
	})
	if count > 0 {
		return nil, v1.ErrorBadRequest("名称已存在")
	}
	return suc.repo.CreateExample(ctx, example)
}

func (suc *ExampleUsecase) UpdateExample(ctx context.Context, example *domain.Example) error {
	count, _ := suc.repo.GetExampleCount(ctx, map[string]interface{}{
		"name":       example.Name,
		"id_deleted": true,
		"id_neq":     example.Id,
	})
	if count > 0 {
		return v1.ErrorBadRequest("名称已存在")
	}
	return suc.repo.UpdateExample(ctx, example)
}

func (suc *ExampleUsecase) DeleteExample(ctx context.Context, example *domain.Example) error {
	count, _ := suc.repo.GetExampleCount(ctx, map[string]interface{}{
		"id": example.Id,
	})
	if count == 0 {
		return v1.ErrorRecordNotFound("数据不存在")
	}
	return suc.repo.DeleteExample(ctx, example)
}

func (suc *ExampleUsecase) RecoverExample(ctx context.Context, example *domain.Example) error {
	count, _ := suc.repo.GetExampleCount(ctx, map[string]interface{}{
		"id":         example.Id,
		"is_deleted": true,
	})
	if count == 0 {
		return v1.ErrorRecordNotFound("数据不存在")
	}
	return suc.repo.RecoverExample(ctx, example)
}
