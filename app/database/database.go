package database

import (
	"fmt"
	"log"
	"time"
	"ushas/config"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	d *gorm.DB

	migrations = &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
)

// Init : Initialize gorm connection.
func Init(isReset bool, models ...interface{}) {
	c := config.GetConfig()
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		c.GetString("db.user"),
		c.GetString("db.password"),
		c.GetString("db.host"),
		c.GetString("db.port"),
		c.GetString("db.database"),
	)
	d, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
