package tool

import (
	"database/sql"
	"sync"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mockInstance *gorm.DB        // 全局唯一的模拟数据库实例
	sqlMock      sqlmock.Sqlmock // 关联的mock控制器
	once         sync.Once       // 确保只初始化一次
)

// GetMysqlMock 返回单例的gorm.DB和sqlmock
func GetMysqlMock() (*gorm.DB, sqlmock.Sqlmock) {
	once.Do(func() {
		// 1. 创建sqlmock实例
		var db *sql.DB
		var err error
		db, sqlMock, err = sqlmock.New()
		if err != nil {
			panic("创建sqlmock失败: " + err.Error())
		}

		// 2. 用mock连接创建gorm.DB
		mockInstance, err = gorm.Open(mysql.New(mysql.Config{
			Conn:                      db,   // 使用mock的sql.DB
			SkipInitializeWithVersion: true, // 跳过版本检查
		}), &gorm.Config{})
		if err != nil {
			panic("创建gorm实例失败: " + err.Error())
		}
	})
	return mockInstance, sqlMock
}
