package domain

import (
	"time"

	"github.com/jinzhu/copier"

	exampleV1 "github.com/ZQCard/kratos-base-layout/api/example/v1"
)

type Example struct {
	Id       int64
	Name     string
	Status   bool
	CreateAt string
	UpdateAt string
}

func (example *Example) ID(id int64) {
	example.Id = id
}

func (example *Example) CreatedAt(timestamp time.Time) {
	example.CreateAt = timestamp.Format("2006-01-02 15:04:05")
}

func (example *Example) UpdatedAt(timestamp time.Time) {
	example.UpdateAt = timestamp.Format("2006-01-02 15:04:05")
}

// Pb 将domain结构体转换为pb结构体
func (example *Example) Pb() *exampleV1.Example {
	pb := &exampleV1.Example{}
	copier.Copy(pb, example)
	return pb
}

// Domain 将pb结构体转换为domain结构体
func (example *Example) Domain(data interface{}) *Example {
	copier.Copy(example, data)
	return example
}

// ToDomainExample 将pb结构体转换为domain包下的Example结构体
func ToDomainExample(data interface{}) *Example {
	example := &Example{}
	copier.Copy(example, data)
	return example
}
