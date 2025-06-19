package config

import (
	"fmt"
	"time"
	"tugasakhir/internal/helper"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	username := viper.GetString("DB_USERNAME")
	password := viper.GetString("DB_PASSWORD")
	host := viper.GetString("DB_HOST")
	port := viper.GetInt("DB_PORT")
	database := viper.GetString("DB_NAME")
	idleConnection := viper.GetInt("DB_POOL_IDLE")
	maxConnection := viper.GetInt("DB_POOL_MAX")
	idleLifeTimeConnection := viper.GetInt("DB_POOL_IDLE_LIFETIME")
	maxLifeTimeConnection := viper.GetInt("DB_POOL_MAX_LIFETIME")

	// fmt.Println(username, password, host, port, database)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             5 * time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	helper.FatalIfErrorWithMessage(err, fmt.Sprintf("Failed to connect database: %s", err))

	connection, err := db.DB()
	helper.FatalIfErrorWithMessage(err, "Failed to setup connection")

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxIdleTime(time.Second * time.Duration(idleLifeTimeConnection))
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}

// migrate -database "mysql://root@tcp(localhost:3306)/academic" -path db/migrations down
