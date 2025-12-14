package database

import (
	"astrodev.online/multimedia-server/internal/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DateTimePrecision = 2
func New(url string) (*gorm.DB,error){
	db,err := gorm.Open(mysql.New(mysql.Config{
     DSN: url,
	 DefaultStringSize: 256,
	 DefaultDatetimePrecision: &DateTimePrecision,
	 DontSupportRenameIndex: true,
	 DontSupportRenameColumn: true,
	 SkipInitializeWithVersion: false,
	}),&gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.AutoMigrate(&models.Upload{})
	return db,err
}