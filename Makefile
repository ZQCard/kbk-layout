GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
PEOJECT ?= "kratos-base-layout"
ServiceUpperName ?= "Example"
ServiceLowerName ?= "example"

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
		   --validate_out=lang=go,paths=source_relative:./api \
		   --go-errors_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

# generate errors proto
.PHONY: errors
errors:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
		   --go-errors_out=paths=source_relative:./api \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

.PHONY: initProject
initProject:
# 初始化项目，将package改为project, 即 module github.com/ZQCard/$(PROJECT)
	@echo "PROJECT=$(PROJECT)"
	@grep  -rl $(PROJECT) ./ | xargs sed -i 's/$(PROJECT)/github.com\/ZQCard\/$(PROJECT)/g'

.PHONY: initNewService
initNewService:
# 复制proto文件
	@mkdir -p ./api/$(ServiceLowerName)/v1/
	@cp ./api/example/v1/example.proto ./api/$(ServiceLowerName)/v1/$(ServiceLowerName).proto
	@cp ./api/example/v1/error_reason.proto ./api/$(ServiceLowerName)/v1/error_reason.proto

# 删除旧文件
	@rm -rf ./api/example
# 生成proto客户端文件
	@ kratos proto client ./api/$(ServiceLowerName)/v1/$(ServiceLowerName).proto
# 替换 Example $(ServiceUpperName), example 为$(ServiceLowerName)并迁移文件
	@grep  -rl Example ./ | xargs sed -i 's/Example/$(ServiceUpperName)/g'
	@grep  -rl example ./ | xargs sed -i 's/example/$(ServiceLowerName)/g'

# 迁移文件
	@ mv ./internal/service/example.go ./internal/service/$(ServiceLowerName).go
	@ mv ./internal/biz/example.go ./internal/biz/$(ServiceLowerName).go
	@ mv ./internal/data/example.go ./internal/data/$(ServiceLowerName).go
	@ mv ./internal/domain/example.go ./internal/domain/$(ServiceLowerName).go
# 拉取引用包
	go mod tidy
	@echo "project start success"
