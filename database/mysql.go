package database

import (
	"fmt"
	"go-tugasku/configs"
	"go-tugasku/models"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var MySQLDB *DBInstance

func MySQLConnect() {
	user := configs.Config("DB_USER")
	pass := configs.Config("DB_PASSWORD")
	host := configs.Config("DB_HOST")
	p := configs.Config("DB_PORT")
	name := configs.Config("DB_NAME")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database!, \n", err)
	}

	log.Println("Database connected!")
	db.Logger = db.Logger.LogMode(logger.Info)

	// Migrate the schema
	db.AutoMigrate(
		&models.Task{},
	)
	log.Println("Database migrated!")

	MySQLDB = &DBInstance{
		DB: db,
	}
}
