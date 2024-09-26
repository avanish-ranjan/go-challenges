package Controllers

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

func GormDB() (*gorm.DB, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbTcp := "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/"
	gormDb, err := gorm.Open(dbDriver, dbUser+":"+dbPassword+dbTcp+dbName+"?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Errorf("gorm Db connection ", err)
		return nil, err
	}
	return gormDb, nil
}
