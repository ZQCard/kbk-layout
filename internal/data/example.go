package data

import (
	"context"
	"net/http"

	"github.com/ZQCard/kratos-base-layout/pkg/utils/timeHelper"

	"github.com/ZQCard/kratos-base-layout/internal/biz"
	"github.com/ZQCard/kratos-base-layout/internal/domain"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type ExampleEntity struct {
	BaseFields
	Name   string `gorm:"type:varchar(255);not null;unique"`
	Status bool   `gorm:"not null"`
}

func (ExampleEntity) TableName() string {
	return "example"
}

type ExampleRepo struct {
	data *Data
	log  *log.Helper
}

func NewExampleRepo(data *Data, logger log.Logger) biz.ExampleRepo {
	repo := &ExampleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/example")),
	}
	return repo
}

// searchParam 搜索条件
func (repo ExampleRepo) searchParam(params map[string]interface{}) *gorm.DB {
	conn := repo.data.db.Model(&ExampleEntity{})
	if ID, ok := params["ID"]; ok && ID.(int64) != 0 {
		conn = conn.Where("ID = ?", ID)
	}
	if name, ok := params["name"]; ok && name.(string) != "" {
		conn = conn.Where("name LIKE ?", "%"+name.(string)+"%")
	}
	// 开始时间
	if start, ok := params["created_at_start"]; ok && start.(string) != "" {
		conn = conn.Where("created_at >= ", start.(string))
	}
	// 结束时间
	if end, ok := params["created_at_end"]; ok && end.(string) != "" {
		conn = conn.Where("created_at <= ", end.(string))
	}

	return conn
}

func (repo ExampleRepo) GetExampleByParams(params map[string]interface{}) (record *ExampleEntity, err error) {
	if len(params) == 0 {
		return &ExampleEntity{}, errors.BadRequest("MISSING_CONDITION", "缺少搜索条件")
	}
	conn := repo.searchParam(params)
	if err = conn.First(record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &ExampleEntity{}, errors.BadRequest("RECORD_NOT_FOUND", "数据不存在")
		}
		return record, errors.InternalServer("SYSTEM_ERROR", err.Error())
	}
	return record, nil
}

func (repo ExampleRepo) CreateExample(ctx context.Context, domain *domain.Example) (*domain.Example, error) {
	entity := &ExampleEntity{}
	entity.ID = domain.ID
	entity.Name = domain.Name
	entity.Name = domain.Name
	if err := repo.data.db.Model(entity).Create(entity).Error; err != nil {
		return nil, errors.New(http.StatusInternalServerError, "SYSTEM_ERROR", err.Error())
	}
	response := toDomainExample(entity)
	return response, nil
}

func (repo ExampleRepo) UpdateExample(ctx context.Context, domain *domain.Example) error {
	// 根据ID查找记录
	record, err := repo.GetExampleByParams(map[string]interface{}{
		"ID": domain.ID,
	})
	if err != nil {
		return err
	}
	// 更新字段
	record.Name = domain.Name
	if err := repo.data.db.Model(&record).Where("ID = ?", record.ID).Save(&record).Error; err != nil {
		return errors.New(http.StatusInternalServerError, "SYSTEM_ERROR", err.Error())
	}

	return nil
}

func (repo ExampleRepo) GetExample(ctx context.Context, params map[string]interface{}) (*domain.Example, error) {
	// 根据ID查找记录
	record, err := repo.GetExampleByParams(params)
	if err != nil {
		return nil, err
	}
	// 返回数据
	response := toDomainExample(record)
	return response, nil
}

func (repo ExampleRepo) ListExample(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*domain.Example, int64, error) {
	list := []*ExampleEntity{}
	conn := repo.searchParam(params)
	err := conn.Scopes(Paginate(page, pageSize)).Find(&list).Error
	if err != nil {
		return nil, 0, errors.New(http.StatusInternalServerError, "SYSTEM_ERROR", err.Error())
	}

	count := int64(0)
	conn.Count(&count)
	rv := make([]*domain.Example, 0, len(list))
	for _, record := range list {
		example := toDomainExample(record)
		rv = append(rv, example)
	}
	return rv, count, nil
}

func (repo ExampleRepo) DeleteExample(ctx context.Context, domain *domain.Example) error {
	// 根据ID查找记录
	record, err := repo.GetExampleByParams(map[string]interface{}{
		"ID": domain.ID,
	})
	if err != nil {
		return err
	}
	if domain.ID != record.ID {
		return errors.New(http.StatusBadRequest, "RECORD_NOT_FOUND", "数据不存在")
	}
	return repo.data.db.Model(&record).Where("ID = ?", domain.ID).UpdateColumn("deleted_at", timeHelper.GetCurrentYMDHIS()).Error
}

func (repo ExampleRepo) RecoverExample(ctx context.Context, ID int64) error {
	if ID == 0 {
		return errors.New(http.StatusBadRequest, "MISSING_CONDITION", "缺少搜索条件")
	}
	err := repo.data.db.Model(ExampleEntity{}).Where("ID = ?", ID).UpdateColumn("deleted_at", "").Error
	if err != nil {
		return errors.New(http.StatusInternalServerError, "SYSTEM_ERROR", err.Error())
	}
	return nil
}

func toDomainExample(example *ExampleEntity) *domain.Example {
	return &domain.Example{
		ID:        example.ID,
		Name:      example.Name,
		CreatedAt: example.CreatedAt,
		UpdatedAt: example.UpdatedAt,
	}
}
