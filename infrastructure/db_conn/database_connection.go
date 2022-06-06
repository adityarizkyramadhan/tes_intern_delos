package db_con

import (
	"fmt"

	"github.com/adityarizkyramadhan/tes_intern_delos/domain"
	"github.com/adityarizkyramadhan/tes_intern_delos/infrastructure/database"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitSupabase(envDb database.DriverSupabase) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s "+
		"password=%s "+
		"host=%s "+
		"TimeZone=Asia/Singapore "+
		"port=%s "+
		"dbname=%s", envDb.User, envDb.Password, envDb.Host, envDb.Port, envDb.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&domain.FarmModel{}, &domain.PondModel{}); err != nil {
		return nil, err
	}
	return db, err
}

func InitMYSql(envDb database.DriverMysql) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", envDb.User, envDb.Password, envDb.Host, envDb.Port, envDb.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&domain.FarmModel{}, &domain.PondModel{}, &domain.CounterEndpoint{}); err != nil {
		return nil, err
	}
	return db, err
}
