package data

import (
	"context"
	"net/http"

	"github.com/ZQCard/kratos-base-layout/internal/biz"
	"github.com/ZQCard/kratos-base-layout/internal/domain"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type ExampleEntity struct {
	BaseFields
	Name   string `gorm:"type:varchar(255);not null;unique;comment:名称"`
	Status bool   `gorm:"not null;comment:状态0冻结1正常"`
}

func (ExampleEntity) TableName() string {
	return "example"
}

type ExampleRepo struct {
	data *Data
	log  *log.Helper
}

func NewExampleRepo(data *Data, logger log.Logger) biz.ExampleRepo {
	r := &ExampleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/example")),
	}
	return r
}

// searchParam 搜索条件
func (r ExampleRepo) searchParam(params map[string]interface{}) *gorm.DB {
	conn := r.data.db.Model(&ExampleEntity{})
	if Id, ok := params["id"]; ok && Id.(int64) != 0 {
		conn = conn.Where("id = ?", Id)
	}
	if name, ok := params["name"]; ok && name.(string) != "" {
		conn = conn.Where("name LIKE ?", "%"+name.(string)+"%")
	}
	// 开始时间
	if start, ok := params["created_at_start"]; ok && start.(string) != "" {
		conn = conn.Where("created_at >= ?", start.(string)+" 00:00:00")
	}
	// 结束时间
	if end, ok := params["created_at_end"]; ok && end.(string) != "" {
		conn = conn.Where("created_at <= ?", end.(string)+" 23:59:59")
	}

	return conn
}

func (r ExampleRepo) GetExampleByParams(params map[string]interface{}) (record *ExampleEntity, err error) {
	if len(params) == 0 {
		return &ExampleEntity{}, errors.BadRequest("MISSING_CONDITION", "缺少搜索条件")
	}
	conn := r.searchParam(params)
	if err = conn.First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &ExampleEntity{}, errors.BadRequest("RECORD_NOT_FOUND", "数据不存在")
		}
		return record, errors.InternalServer("SYSTEM ERROR", err.Error())
	}
	return record, nil
}

func (r ExampleRepo) CreateExample(ctx context.Context, domain *domain.Example) (*domain.Example, error) {
	entity := &ExampleEntity{}
	entity.Id = domain.Id
	entity.Name = domain.Name
	entity.Status = domain.Status
	if err := r.data.db.Model(entity).Create(entity).Error; err != nil {
		r.log.Errorf("CreateExample error: %v", err)
		return nil, errors.New(http.StatusInternalServerError, "SYSTEM ERROR", err.Error())
	}
	response := toDomainExample(entity)
	return response, nil
}

func (r ExampleRepo) UpdateExample(ctx context.Context, domain *domain.Example) error {
	// 根据Id查找记录
	record, err := r.GetExampleByParams(map[string]interface{}{
		"id": domain.Id,
	})
	if err != nil {
		return err
	}
	// 更新字段
	record.Name = domain.Name
	record.Status = domain.Status
	if err := r.data.db.Model(&record).Where("id = ?", record.Id).Save(&record).Error; err != nil {
		r.log.Errorf("UpdateExample error: %v", err)
		return errors.New(http.StatusInternalServerError, "SYSTEM ERROR", err.Error())
	}

	return nil
}

func (r ExampleRepo) GetExample(ctx context.Context, params map[string]interface{}) (*domain.Example, error) {
	// 根据Id查找记录
	record, err := r.GetExampleByParams(params)
	if err != nil {
		return nil, err
	}
	// 返回数据
	response := toDomainExample(record)
	return response, nil
}

func (r ExampleRepo) ListExample(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*domain.Example, int64, error) {
	list := []*ExampleEntity{}
	conn := r.searchParam(params)
	err := conn.Scopes(Paginate(page, pageSize)).Find(&list).Error
	if err != nil {
		r.log.Errorf("ListExample error: %v", err)
		return nil, 0, errors.New(http.StatusInternalServerError, "SYSTEM ERROR", err.Error())
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

func (r ExampleRepo) DeleteExample(ctx context.Context, domain *domain.Example) error {
	// 根据Id查找记录
	record, err := r.GetExampleByParams(map[string]interface{}{
		"id": domain.Id,
	})
	if err != nil {
		return err
	}
	if domain.Id != record.Id {
		return errors.BadRequest("RECORD_NOT_FOUND", "数据不存在")
	}
	if err := r.data.db.Where("Id = ?", domain.Id).Delete(&ExampleEntity{}).Error; err != nil {
		r.log.Errorf("DeleteExample error: %v", err)
		return errors.InternalServer("SYSTEM ERROR", err.Error())
	}
	return nil
}

func (r ExampleRepo) RecoverExample(ctx context.Context, domain *domain.Example) error {
	if domain.Id == 0 {
		return errors.BadRequest("MISSING_CONDITION", "缺少搜索条件")
	}
	if err := r.data.db.Model(ExampleEntity{}).Where("Id = ?", domain.Id).UpdateColumn("deleted_at", "").Error; err != nil {
		r.log.Errorf("RecoverExample error: %v", err)
		return errors.New(http.StatusInternalServerError, "SYSTEM ERROR", err.Error())
	}
	return nil
}

func toDomainExample(example *ExampleEntity) *domain.Example {
	if example == nil {
		return &domain.Example{}
	}
	return &domain.Example{
		Id:        example.Id,
		Name:      example.Name,
		CreatedAt: example.CreatedAt,
		UpdatedAt: example.UpdatedAt,
	}
}
