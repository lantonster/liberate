package data

import (
	"fmt"

	"github.com/lantonster/liberate/internal/config"
	"github.com/lantonster/liberate/pkg/orm"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB creates a new database connection using the provided configuration.
// It returns a *gorm.DB instance and an error if the connection fails.
//
// Parameters:
//   - cfg: The configuration containing MySQL connection details
//
// Returns:
//   - *gorm.DB: The database connection instance
//   - error: Any error encountered during connection
func NewDB(cfg *config.Config) *gorm.DB {
	// 构建 MySQL 连接 DSN 字符串，包含用户名、密码、主机地址、端口和数据库名称
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.Username,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.Database,
	)

	// 使用 gorm.Open 建立数据库连接，如果失败则 panic
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("db connected")

	// 设置默认数据库连接
	orm.SetDefault(db)

	return db
}
