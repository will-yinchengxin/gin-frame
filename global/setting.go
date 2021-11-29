package global

import (
	"frame/pkg/logger"
	"frame/pkg/setting"
	"frame/pkg/tracer"
	"frame/pkg/validator"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/natefinch/lumberjack"
	"github.com/opentracing/opentracing-go"
	"log"
	"time"
)

var (
	// 基础配置
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	TracerSetting   *setting.JaegerSetting
	RedisSetting    map[string]*setting.RedisSetting
	RocketSetting   *setting.RocketSetting
	// 上传文件配置
	UploadFileSetting *setting.UploadFile
	// 数据库
	DBEngine *gorm.DB
	// redis 连接池
	Redis map[string]*redis.Pool
	// 日志连接
	Logger *logger.Logger
	// 验证器
	ReqValidator *validator.ValidatorX
	// jeager
	Tracer opentracing.Tracer
	// rocketMQ
	RocketConsumer rocketmq.PushConsumer
	RocketProducer rocketmq.Producer
)

func SetupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Redis", &RedisSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("UploadFile", &UploadFileSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Jaeger", &TracerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("RocketMQ", &RocketSetting)
	if err != nil {
		return err
	}

	ServerSetting.ReadTimeout = time.Second
	ServerSetting.WriteTimeout = time.Second

	return nil
}

func SetupLogger() (err error) {
	// log.LstdFlags 标准记录器的初始值
	// 这里使用了lumberjack作为日志库的io.Writer,并将日志文件所允许的最大占用空间设置未600M,日志文件的最大生成周期10天,日志格为本地时间

	// 这里使用的是windows,所以给出了绝对的格式路径
	// fileName := "C:\\Users\\Administrator\\Desktop\\go\\code\\frame\\storage\\logs\\app.log"
	fileName := AppSetting.LogSavePath + AppSetting.LogFileName + AppSetting.LogFileExt
	Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    600,  // 兆字节
		MaxBackups: 10,   //
		MaxAge:     10,   // days
		Compress:   true, // 默认禁用
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func SetValidator() {
	ReqValidator = validator.NewValidator()
	return
}

func SetTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(TracerSetting.Name, TracerSetting.Host)
	if err != nil {
		return err
	}
	Tracer = jaegerTracer
	return nil
}

// 建立 redis 连接池
func SetRedis() (err error) {
	if len(RedisSetting) < 1 {
		panic("init redis pool config failed, redis config not found")
	}
	Redis = make(map[string]*redis.Pool)
	for name, val := range RedisSetting {
		if rdb, err := initRedisPool(val); rdb != nil {
			if err != nil {
				return err
			}
			Redis[name] = rdb
		}
	}
	return
}

// 初始化 rocketMQ
func SetRocketMQ() (err error) {
	RocketProducer, err = rocketmq.NewProducer(
		producer.WithNameServer(RocketSetting.Host),
		producer.WithRetry(RocketSetting.Retry),
		producer.WithGroupName(RocketSetting.GroupName),
	)
	if err != nil {
		Logger.Fatalf("NewProducer fail, cause %s", err)
		return err
	}

	// 生产者
	err = RocketProducer.Start()
	if err != nil {
		Logger.Fatalf("Producer Start fail, cause %s", err)
		return err
	}

	// 消费者
	RocketConsumer, err = rocketmq.NewPushConsumer(
		consumer.WithNameServer(RocketSetting.Host),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName(RocketSetting.GroupName),
	)
	if err != nil {
		Logger.Fatalf("NewPushConsumer fail, cause %s", err)
		return err
	}

	return
}