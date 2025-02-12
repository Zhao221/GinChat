package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

const (
	PublishKey = "websocket"
)

func InitConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func InitMysql() error {
	// 自定义日志模版 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: time.Second, // 慢SQL阀值
			LogLevel:      logger.Info, // 级别
			Colorful:      true,        // 彩色
		},
	)
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: newLogger})
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func InitRedis() error {
	// 初始化redis客户端
	redisClient := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.db"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	Redis = redisClient
	return nil
}

// Publish 发送消息到redis
func Publish(ctx context.Context, channel string, message interface{}) error {
	var err error
	err = Redis.Publish(ctx, channel, message).Err()
	if err != nil {
		return err
	}
	return nil
}

// Subscribe 订阅redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Redis.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		return "", err
	}
	return msg.Payload, nil
}
