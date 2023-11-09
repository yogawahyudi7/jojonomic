package utils

import (
	"fmt"
	"time"

	"github.com/yogawahyudi7/jojonomic/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func PostgreSQL(config *config.ServerConfig) *gorm.DB {

	set := config.Database.PostgreSQL

	sslmode := "disable"
	timeZone := "Asia/Jakarta"
	user := set.Username
	password := set.Password
	host := set.Host
	port := set.Port
	dbName := set.Name

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", host, user, password, dbName, port, sslmode, timeZone)

	fmt.Println("--DNS CONNECTION--")
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // open connection

	if err != nil {
		panic(err)
	}

	db.Use(dbresolver.Register(dbresolver.Config{}).SetMaxIdleConns(10).SetMaxOpenConns(100).SetConnMaxLifetime(time.Hour))

	return db
}
