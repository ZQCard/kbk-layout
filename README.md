# kratos-base-layout
方便快速启动服务， 统一项目模板

# 1.创建项目指定模板
假设项目目录为data-center

```
kratos new $(projectName) -r git@github.com:ZQCard/kratos-base-layout.git
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
以上操作将生成目录为 data-center, 包含服务名称为StoreService的crud服务。

log: kratos/v2/log

orm: gorm

doc: swagger

proto-validate:protoc-gen-validate

db: mysql

cache:redis

pkg/utils: 常用函数包

# 其他注意事项

0. 环境设置为dev，则为开发环境，可以根据该变量设定规则， 如 ent会启用debug模式,打印sql

1. 整形数值类型均用int64， proto转json的时候由于精度问题，会将int64类型转为string类型

2. http://localhost:8000/q/swagger-ui/ 查看接口文档

3. 使用protoc-gen-validate进行proto接口参数验证


