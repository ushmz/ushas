package database

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"ushas/config"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	d *gorm.DB

	migrations = &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}

	gormLogger logger.Interface

	devLogger  = log.New(os.Stderr, "[gorm] ", log.LstdFlags|log.Lshortfile)
	prodLogger = log.New(io.Discard, "[gorm] ", log.LstdFlags)

	devConfig = logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Error,
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	}

	prodConfig = logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Silent,
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	}
)

// Init : Initialize gorm connection.
func Init(isReset bool, models ...interface{}) {
	c := config.GetConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		c.GetString("db.user"),
		c.GetString("db.password"),
		c.GetString("db.host"),
		c.GetString("db.port"),
		c.GetString("db.database"),
	)

	if env := c.GetString("env"); env == "dev" {
		gormLogger = logger.New(devLogger, devConfig)
	} else {
		gormLogger = logger.New(prodLogger, prodConfig)
	}

	// To avoid shadowing (`d`), declare `err` variable with nil value
	// and do not use `:=` operator
	var err error
	d, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		// Failed to establish gorm session.
		panic(err)
	}

	db, err := d.DB()
	if err != nil {
		// Invalid DB configuration.
		panic(err)
	}

	// Wait until connect to DB
	for {
		if err := db.Ping(); err != nil {
			log.Println("DB is not ready. Retry connecting...")
			time.Sleep(1 * time.Second)
			continue
		}
		log.Println("Success to connect DB")
		break
	}

	// Auto migration
	if c.GetString("env") == "prod" {
		d.AutoMigrate(models...)
	}
}

// GetDB : Get new gorm connection.
func GetDB() *gorm.DB {
	return d
}
