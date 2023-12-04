package databases

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"errors"
	"go-practice/backendApi/models"
)

type DbInstance struct{
	DB *gorm.DB 
}

func GetDBInstance ()(DbInstance){
	instance, _ := CockroachInit()
	return instance
}

func CockroachInit()(DbInstance, error) {
	db, err := CockroachConnect()
	if err != "" {
		return DbInstance{nil}, errors.New(err)
	}
	dbInstance := SchemaMigrations(db)
	return dbInstance, nil
}

func SchemaMigrations(db *gorm.DB) (DbInstance){
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.BookAttrs{})
	return DbInstance{DB:db}
}

func CockroachConnect()(*gorm.DB, string) {
	dsn:="URL"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err.Error()
}