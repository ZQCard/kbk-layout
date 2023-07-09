package data

import (
	"context"

	"github.com/ZQCard/kbk-layout/internal/biz"
	"github.com/ZQCard/kbk-layout/internal/domain"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	exampleV1 "github.com/ZQCard/kbk-layout/api/example/v1"
)

type ExampleEntity struct {
	BaseFields
	Name   string `gorm:"type:varchar(255);not null;comment:名称"`
	Status bool   `gorm:"not null;comment:状态0冻结1正常"`
	Domain string `gorm:"type:varchar(255);not null;comment:域"`
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
func (r ExampleRepo) searchParam(ctx context.Context, params map[string]interface{}) *gorm.DB {
	conn := r.data.db.Model(&ExampleEntity{})
	if Id, ok := params["id"]; ok && Id.(int64) != 0 {
		conn = conn.Where("id = ?", Id)
	}
	if v, ok := params["name"]; ok && v.(string) != "" {
		conn = conn.Where("name LIKE ?", "%"+v.(string)+"%")
	}
	if v, ok := params["status"]; ok && v != nil {
		conn = conn.Where("status = ?", v)
	}
	// 开始时间
	if v, ok := params["created_at_start"]; ok && v.(string) != "" {
		conn = conn.Where("created_at >= ?", v.(string)+" 00:00:00")
	}
	// 结束时间
	if v, ok := params["created_at_end"]; ok && v.(string) != "" {
		conn = conn.Where("created_at <= ?", v.(string)+" 23:59:59")
	}
	if v, ok := params["is_deleted"]; ok {
		if v.(bool) {
			conn = conn.Unscoped()
		}
	}
	conn = getDbWithDomain(ctx, conn)
	return conn
}

func (r ExampleRepo) ListExample(ctx context.Context, page, pageSize int64, params map[string]interface{}) ([]*domain.Example, error) {
	list := []*ExampleEntity{}
	conn := r.searchParam(ctx, params)
	err := conn.Scopes(Paginate(page, pageSize)).Find(&list).Error
	if err != nil {
		return nil, exampleV1.ErrorSystemError("获取列表失败").WithCause(err)
	}

	rv := make([]*domain.Example, 0, len(list))
	for _, record := range list {
		rv = append(rv, domain.ToDomainExample(record))
	}
	return rv, nil
}
func (r ExampleRepo) GetExampleCount(ctx context.Context, params map[string]interface{}) (count int64, err error) {
	if len(params) == 0 {
		return 0, exampleV1.ErrorBadRequest("参数不得为空")
	}
	r.searchParam(ctx, params).Count(&count)
	return count, nil
}

func (r ExampleRepo) GetExampleByParams(ctx context.Context, params map[string]interface{}) (record *ExampleEntity, err error) {
	if len(params) == 0 {
		return &ExampleEntity{}, exampleV1.ErrorBadRequest("参数不得为空")
	}
	conn := r.searchParam(ctx, params)
	if err = conn.First(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &ExampleEntity{}, exampleV1.ErrorRecordNotFound("数据不存在")
		}
		return record, exampleV1.ErrorSystemError("查询记录失败").WithCause(err)
	}
	return record, nil
}

func (r ExampleRepo) CreateExample(ctx context.Context, example *domain.Example) (*domain.Example, error) {
	entity := &ExampleEntity{}
	entity.Id = example.Id
	entity.Name = example.Name
	entity.Status = example.Status
	entity.Domain = example.Domain
	if err := r.data.db.Model(entity).Create(entity).Error; err != nil {
		return nil, exampleV1.ErrorSystemError("创建失败").WithCause(err)
	}
	return domain.ToDomainExample(entity), nil
}

func (r ExampleRepo) UpdateExample(ctx context.Context, domain *domain.Example) error {
	// 根据Id查找记录
	record, err := r.GetExampleByParams(ctx, map[string]interface{}{
		"id": domain.Id,
	})
	if err != nil {
		return err
	}
	// 更新字段
	record.Name = domain.Name
	record.Status = domain.Status
	record.Domain = domain.Domain
	if err := r.data.db.Model(&record).Where("id = ?", record.Id).Save(&record).Error; err != nil {
		return exampleV1.ErrorSystemError("更新失败").WithCause(err)
	}

	return nil
}

func (r ExampleRepo) GetExample(ctx context.Context, params map[string]interface{}) (*domain.Example, error) {
	// 根据Id查找记录
	record, err := r.GetExampleByParams(ctx, params)
	if err != nil {
		return nil, err
	}
	// 返回数据
	response := domain.ToDomainExample(record)
	return response, nil
}

func (r ExampleRepo) DeleteExample(ctx context.Context, domain *domain.Example) error {
	// 根据Id查找记录
	record, err := r.GetExampleByParams(ctx, map[string]interface{}{
		"id": domain.Id,
	})
	if err != nil {
		return err
	}
	if domain.Id != record.Id {
		return exampleV1.ErrorRecordNotFound("数据不存在")
	}
	if err := r.data.db.Where("id = ?", domain.Id).Delete(&ExampleEntity{}).Error; err != nil {
		return exampleV1.ErrorSystemError("删除数据失败").WithCause(err)
	}
	return nil
}

func (r ExampleRepo) RecoverExample(ctx context.Context, domain *domain.Example) error {
	if domain.Id == 0 {
		return exampleV1.ErrorBadRequest("缺少搜索条件")
	}
	// 根据Id查找记录
	record, err := r.GetExampleByParams(ctx, map[string]interface{}{
		"id":         domain.Id,
		"is_deleted": true,
	})
	if err != nil {
		return err
	}
	if domain.Id != record.Id {
		return exampleV1.ErrorRecordNotFound("数据不存在")
	}
	if err := r.data.db.Model(ExampleEntity{}).Unscoped().Where("id = ?", domain.Id).UpdateColumn("deleted_at", "").Error; err != nil {
		return exampleV1.ErrorSystemError("恢复数据失败").WithCause(err)
	}
	return nil
}
