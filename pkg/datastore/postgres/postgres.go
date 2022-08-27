package postgres

import (
	"fmt"
	"github.com/bitwyre/template-go/pkg/datastore/postgres/entity"
	"log"

	"github.com/bitwyre/template-go/pkg/lib"
	pgDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ClientInstance struct {
	Db *gorm.DB
}

func NewClient() *ClientInstance {
	var env = lib.AppConfig.App
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta`, env.PgHost, env.PgUser, env.PgPassword, env.PgDB, env.PgPort)

	db, err := gorm.Open(pgDriver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln("🔴  Database Error:" + err.Error())
		return nil
	}

	log.Println("🚀 Postgres database connected")

	return &ClientInstance{Db: db}
}

func (client *ClientInstance) AutoMigrate() {
	if lib.AppConfig.App.Env != "prod" {
		if lib.AppConfig.App.Migration == "y" {
			err := client.Db.AutoMigrate(
				&entity.AuthToken{},
				&entity.UserAuth{},
			)
			if err != nil {
				log.Fatalln("🔴 Database Migration Failed", err)
			}
			log.Println("🟢 Database Migration Success")
			return
		}

		log.Println("🟡 Database Migration Skipped")
	}
}
