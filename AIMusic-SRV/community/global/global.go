package global

import (
	"community/config"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var (
	DB     *gorm.DB
	RDB    *redis.Client
	Config = &config.AIMusicConfig{}
)

func configureInit() {
	v := viper.New()
	v.SetConfigFile("config.yml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(Config); err != nil {
		panic(err)
	}
}

func init() {
	configureInit()
	redisInit()
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/AIMusic?charset=utf8mb4&parseTime=True&loc=Local",
		Config.MySQL.User, Config.MySQL.Password, Config.MySQL.Host, Config.MySQL.Port)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

func redisInit() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", Config.Redis.Host, Config.Redis.Port),
		Password: Config.Redis.Password, // 没有密码，默认值
		DB:       0,                     // 默认DB 0
	})
}
