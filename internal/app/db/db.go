package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myjsonrpcv2/internal/app/conf"
	"myjsonrpcv2/internal/app/constant"
	"myjsonrpcv2/internal/app/model/entity"
	"net/url"
)

var mysqlDb *gorm.DB

func InitMysqlDB(cfg conf.Mysql) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
		//encodeTimeZone(cfg.TimeZone),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("start mysql client failed, db:%s, err:%s", cfg.Database, err.Error())
	}
	mysqlDb = db
}

func GetDb() *gorm.DB {
	return mysqlDb
}

func encodeTimeZone(timezone string) string {
	if timezone == "" {
		timezone = constant.DefaultTimezone
	}

	return url.QueryEscape(fmt.Sprintf("'%s'", timezone))
}

func CreateTable() {
	err := GetDb().AutoMigrate(
		&entity.User{},
	)
	if err != nil {
		logrus.Error(err)
	}
}
