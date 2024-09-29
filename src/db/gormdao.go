package db

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // init mysql driver
	"github.com/jinzhu/gorm"
)

var (
	db       *gorm.DB
	interval = 60
	dialect  = "mysql"
	source   = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		"root",
		"123456",
		"docker.for.mac.localhost",
		3306,
		"eCommerceService",
		"charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci",
	)
)

func init() {
	go connectionPool()
}

// DB return database instance
func DB() *gorm.DB {
	if db == nil {
		fmt.Println(dialect, source)
		connect(dialect, source)
	}
	return db
}

func connectionPool() {
	for {
		if err := DB().DB().Ping(); err != nil {
			connect(dialect, source)
			fmt.Sprintf("Database connection lost, reconnecting...")
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func connect(dialect string, source string) {
	conn, err := gorm.Open(dialect, source)

	if err != nil {
		fmt.Println("Error connecting to database: %v", err)
		os.Exit(1)
	} else {
		conn.SingularTable(true)
		conn.BlockGlobalUpdate(true)
		conn.Exec("SET wait_timeout=300")

		db = conn
	}
}
