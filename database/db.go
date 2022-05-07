package database

import (
	"crud/model"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Pool *gorm.DB

func ConnectDatabase() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	dsn := "root:@tcp(172.17.0.1:33060)/company?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// NamingStrategy: schema.NamingStrategy{
		// 	SingularTable: true,
		// },
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	DB.AutoMigrate(&model.Employee{})
	Pool = DB
	// Migrate the schema
	fmt.Println("DB Connect success!\n")
}

// cfg := mysql.Config{
// 	User:   "root",
// 	Net:    "tcp",
// 	Addr:   "127.0.0.1:33060",
// 	DBName: "Company",
// }
// // Get a database handle.
// var err error
// db, err = sql.Open("mysql", cfg.FormatDSN())
// if err != nil {
// 	log.Fatal(err)
// }
// pingErr := db.Ping()
// if pingErr != nil {
// 	log.Fatal(pingErr)
// }
