package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	etcdclient "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ZQCard/kratos-base-layout/internal/conf"
	"github.com/ZQCard/kratos-base-layout/pkg/middleware/requestInfo"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDiscovery,
	NewRegistrar,
	NewRedisCmd,
	NewMysqlCmd,
	NewRedisClient,
	NewExampleRepo)

// Data .
type Data struct {
	cfg    *conf.Bootstrap
	logger *log.Helper
	db     *gorm.DB
	rdb    *redis.Client
}

func NewData(cfg *conf.Bootstrap, db *gorm.DB, redisCli *redis.Client, logger log.Logger) (*Data, func(), error) {
	logs := log.NewHelper(log.With(logger, "module", "kratos-base-layout/data"))
	cleanup := func() {
		logs.Info("closing the data resources")
	}

	return &Data{
		logger: logs,
		cfg:    cfg,
		db:     db,
		rdb:    redisCli,
	}, cleanup, nil
}

func NewRedisClient(conf *conf.Data) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
		DB:           int(conf.Redis.Db),
	})
	err := client.Ping().Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}

func getDomain(ctx context.Context) string {
	domain := ctx.Value(requestInfo.DomainKey)
	return domain.(string)
}

func getDbWithDomain(ctx context.Context, db *gorm.DB) *gorm.DB {
	domain := ctx.Value(requestInfo.DomainKey)
	if domain != nil {
		db = db.Where("domain = ?", domain)
	}
	return db
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	point := conf.Etcd.Address
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{point},
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	point := conf.Etcd.Address
	client, err := etcdclient.New(etcdclient.Config{
		Endpoints: []string{point},
	})
	if err != nil {
		panic(err)
	}
	r := etcd.New(client)
	return r
}

func NewRedisCmd(conf *conf.Data, logger log.Logger) redis.Cmdable {
	logs := log.NewHelper(log.With(logger, "module", "serviceName/data/redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	err := client.Ping().Err()
	if err != nil {
		logs.Fatalf("redis connect error: %v", err)
	}
	return client
}

func NewMysqlCmd(conf *conf.Bootstrap, logger log.Logger) *gorm.DB {
	logs := log.NewHelper(log.With(logger, "module", "serviceName/data/mysql"))
	db, err := gorm.Open(mysql.Open(conf.Data.Database.Source), &gorm.Config{})
	if err != nil {
		logs.Fatalf("mysql connect error: %v", err)
	}
	// 如果是开发环境 打印sql
	if conf.Env.Mode == "dev" {
		db = db.Debug()
	}
	// 数据迁移
	db.AutoMigrate(&ExampleEntity{})
	return db
}
