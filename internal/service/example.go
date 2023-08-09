package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	exampleV1 "github.com/ZQCard/kbk-layout/api/example/v1"
	"github.com/ZQCard/kbk-layout/internal/biz"
	"github.com/ZQCard/kbk-layout/internal/domain"
)

type ExampleService struct {
	exampleV1.UnimplementedExampleServiceServer
	exampleUsecase *biz.ExampleUsecase
	log            *log.Helper
}

func NewExampleService(exampleUsecase *biz.ExampleUsecase, logger log.Logger) *ExampleService {
	return &ExampleService{
		log:            log.NewHelper(log.With(logger, "module", "service/ExampleService")),
		exampleUsecase: exampleUsecase,
	}
}

func (s *ExampleService) GetExampleList(ctx context.Context, req *exampleV1.GetExampleListReq) (*exampleV1.GetExampleListPageRes, error) {
	params := make(map[string]interface{})
	params["name"] = req.Name
	if req.Status != nil {
		params["status"] = req.Status.Value
	}
	if req.IsDeleted != nil {
		params["is_deleted"] = req.IsDeleted.Value
	}
	params["created_at_start"] = req.CreatedAtStart
	params["created_at_end"] = req.CreatedAtEnd
	list, count, err := s.exampleUsecase.ListExample(ctx, req.Page, req.PageSize, params)
	if err != nil {
		return nil, err
	}
	res := &exampleV1.GetExampleListPageRes{}
	res.Total = int64(count)
	for _, v := range list {
		res.List = append(res.List, v.Pb())
	}
	return res, nil
}

func (s *ExampleService) GetExample(ctx context.Context, req *exampleV1.ExampleIdReq) (*exampleV1.Example, error) {
	res, err := s.exampleUsecase.GetExample(ctx, domain.ToDomainExample(req))
	if err != nil {
		return nil, err
	}
	return res.Pb(), nil
}

func (s *ExampleService) CreateExample(ctx context.Context, req *exampleV1.CreateExampleReq) (*exampleV1.Example, error) {
	res, err := s.exampleUsecase.CreateExample(ctx, domain.ToDomainExample(req))
	if err != nil {
		return nil, err
	}
	return res.Pb(), nil
}

func (s *ExampleService) UpdateExample(ctx context.Context, req *exampleV1.UpdateExampleReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.exampleUsecase.UpdateExample(ctx, domain.ToDomainExample(req))
}

func (s *ExampleService) DeleteExample(ctx context.Context, req *exampleV1.ExampleIdReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.exampleUsecase.DeleteExample(ctx, domain.ToDomainExample(req))
}

func (s *ExampleService) RecoverExample(ctx context.Context, req *exampleV1.ExampleIdReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.exampleUsecase.RecoverExample(ctx, domain.ToDomainExample(req))
}
