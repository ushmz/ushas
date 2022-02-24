package database

import (
	"fmt"
	"ushas/config"

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
		panic(err)
	}

	// Wait until connect to DB
	// for {
	// 	err = d.Ping()
	// 	if err != nil {
	// 		log.Println("DB is not ready. Retry connecting...")
	// 		time.Sleep(1 * time.Second)
	// 		continue
	// 	}
	// 	log.Println("Success to connect DB")
	// 	break
	// }

	// Migration
	if c.GetString("env") == "prod" {
		d.AutoMigrate(models...)
	}

}

func GetDB() *gorm.DB {
	return d
}
