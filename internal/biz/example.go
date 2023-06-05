package biz

import (
	"context"

	"github.com/ZQCard/kratos-base-layout/internal/domain"
	"github.com/go-kratos/kratos/v2/log"
)

type ExampleRepo interface {
	ListExample(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*domain.Example, int64, error)
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
	return suc.repo.ListExample(ctx, page, pageSize, params)
}

func (suc *ExampleUsecase) GetExample(ctx context.Context, example *domain.Example) (*domain.Example, error) {
	params := make(map[string]interface{})
	params["id"] = example.Id
	return suc.repo.GetExample(ctx, params)
}

func (suc *ExampleUsecase) CreateExample(ctx context.Context, example *domain.Example) (*domain.Example, error) {
	return suc.repo.CreateExample(ctx, example)
}

func (suc *ExampleUsecase) UpdateExample(ctx context.Context, example *domain.Example) error {
	return suc.repo.UpdateExample(ctx, example)
}

func (suc *ExampleUsecase) DeleteExample(ctx context.Context, example *domain.Example) error {
	return suc.repo.DeleteExample(ctx, example)
}

func (suc *ExampleUsecase) RecoverExample(ctx context.Context, example *domain.Example) error {
	return suc.repo.RecoverExample(ctx, example)
}
