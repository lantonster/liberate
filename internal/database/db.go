package database

import (
	"context"

	"github.com/lantonster/liberate/internal/config"
	"github.com/lantonster/liberate/pkg/color"
	"github.com/lantonster/liberate/pkg/database"
	"github.com/lantonster/liberate/pkg/log"
	"github.com/lantonster/liberate/pkg/orm"

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
func NewDB(conf *config.Config) *gorm.DB {
	db := database.NewDB(conf.MySQL, log.NewGormLogger(conf.MySQL.Debug))

	log.WithContext(context.Background()).Info(color.Green.Sprint("database connected"))

	// 设置默认数据库连接
	orm.SetDefault(db)

	return db
}
