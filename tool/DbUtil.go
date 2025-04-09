package tool

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

/*
定义数据库连接参数
*/
func InitDb() error {
	dsn, e := getDatabaseConnection()
	if e != nil {
		return e
	}
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	Db = d
	return nil
}

func getDatabaseConnection() (string, error) {
	connection, err := ConfigReadDatabaseConnection()
	if err != nil {
		return "", err
	}
	return connection, nil
}
