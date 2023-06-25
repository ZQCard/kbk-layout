# kbk-layout
方便快速启动服务， 统一项目模板，三键生成微服务

# 1.创建项目指定模板
假设项目目录为projectName

```
kratos new $(projectName) -r git@github.com:ZQCard/kbk-layout.git
```
# 2.初始化项目package
此处的projectName的值需要与第一步保持一致
```
make initProject PROJECT=$(projectName)
```
# 3.初始化项目服务
```
make initNewService ServiceUpperName=Store ServiceLowerName=store
```

# 总结
以上操作将生成目录为 projectName, 包含服务名称为StoreService的crud服务。

* db: mysql, orm: gorm
* cache:redis
* trace:jaeger
* registry && discovery: etcd
* log: kratos/v2/log json格式
* validate: protoc-gen-validate


pkg/utils: 常用函数包

# 其他注意事项

0. 环境设置为dev，则为开发环境，可以根据该变量设定规则， 如 debug模式,打印sql

1. 整形数值类型均用int64， proto转json的时候由于精度问题，会将int64类型转为string类型

2. 接口文档在openapi.yaml中
3. 本地开发以及docker部署文件为configs/config-dev.yaml,可在Dockerfile中自行更改



