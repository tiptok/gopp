package pg

import (
	"fmt"
	"github.com/tiptok/gopp/orm/pgorm/pg/models"
	"github.com/tiptok/gopp/pkg/constant"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable TimeZone=Asia/Shanghai",
		constant.POSTGRESQL_HOST,
		constant.POSTGRESQL_PORT,
		constant.POSTGRESQL_USER,
		constant.POSTGRESQL_DB_NAME,
		constant.POSTGRESQL_PASSWORD,
	)), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db

	if !constant.DISABLE_CREATE_TABLE {
		for _, model := range []interface{}{
			(*models.Users)(nil),
			(*models.SysMessageConsume)(nil),
			(*models.SysMessageProduce)(nil),
		} {
			err := db.AutoMigrate(model)
			if err != nil {
				panic(err)
			}
		}
	}
}
