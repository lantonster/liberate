package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQL struct {
	Host     string `mapstructure:"host" default:"localhost"`
	Port     int    `mapstructure:"port" default:"3306"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Debug    bool   `mapstructure:"debug" default:"false"`
}

func NewDB(conf MySQL, logger logger.Interface) *gorm.DB {
	// 构建 MySQL 连接 DSN 字符串，包含用户名、密码、主机地址、端口和数据库名称
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)

	// 使用 gorm.Open 建立数据库连接，如果失败则 panic
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger})
	if err != nil {
		panic(err)
	}

	return db
}
