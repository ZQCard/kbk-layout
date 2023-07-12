package domain

import (
	"reflect"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	exampleV1 "github.com/ZQCard/kbk-layout/api/example/v1"
)

type Example struct {
	Id        int64
	Domain    string
	Name      string
	Status    bool
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func (example *Example) ID(id int64) {
	example.Id = id
}

// Pb 将domain结构体转换为pb结构体
func (example *Example) Pb() *exampleV1.Example {
	pb := &exampleV1.Example{}
	copier.Copy(pb, example)
	return pb
}

// ToDomainExample 将pb结构体转换为domain包下的Example结构体
func ToDomainExample(data interface{}) *Example {
	example := &Example{}
	if data == nil {
		return example
	}
	copier.Copy(example, data)

	exampleEntity := reflect.ValueOf(data)
	if exampleEntity.Kind() == reflect.Ptr {
		exampleEntity = exampleEntity.Elem()
	}
	if exampleEntity.FieldByName("CreatedAt").IsValid() {
		example.CreatedAt = exampleEntity.FieldByName("CreatedAt").Interface().(time.Time).Format("2006-01-02 15:04:05")
	}
	if exampleEntity.FieldByName("UpdatedAt").IsValid() {
		example.UpdatedAt = exampleEntity.FieldByName("UpdatedAt").Interface().(time.Time).Format("2006-01-02 15:04:05")
	}
	if exampleEntity.FieldByName("DeletedAt").IsValid() {
		deletedAt := exampleEntity.FieldByName("DeletedAt").Interface().(gorm.DeletedAt)
		if deletedAt.Valid {
			example.DeletedAt = deletedAt.Time.Format("2006-01-02 15:04:05")
		} else {
			example.DeletedAt = ""
		}
	}
	return example
}
