package models

import (
	"MonitorProject/tool"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestGetMonitor_Success(t *testing.T) {
	mockDb, mock := tool.GetMysqlMock()

	// 替换全局的 tool.Db 为 mock 实例
	originalDb := tool.Db
	tool.Db = mockDb
	defer func() { tool.Db = originalDb }() // 测试结束后恢复原 DB

	// 2. 定义预期数据
	expectedTarget := MonitorTarget{
		Id:        1,
		Ip:        "192.168.1.1",
		Domain:    "example.com",
		IsDeleted: 0,
	}

	// 3. 设置 Mock 预期, 预期执行的 SQL 查询
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `monitor_target` WHERE id = ? and is_deleted = ?")).
		// 匹配查询参数
		WithArgs(1, 0).
		// 返回模拟数据行
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "ip", "domain", "is_deleted"}).
				AddRow(expectedTarget.Id, expectedTarget.Ip,
					expectedTarget.Domain, expectedTarget.IsDeleted),
		)

	// 4. 调用被测方法
	var target MonitorTarget
	result := tool.Db.Debug().
		Table("monitor_target").
		Where("id = ? and is_deleted = ?", 1, 0).
		Find(&target)

	// 5. 断言结果
	assert.NoError(t, result.Error)                       // 验证没有错误
	assert.Equal(t, expectedTarget.Id, target.Id)         // 验证返回的 ID 符合预期
	assert.Equal(t, expectedTarget.Ip, target.Ip)         // 验证 IP 字段
	assert.Equal(t, expectedTarget.Domain, target.Domain) // 验证 Domain 字段

	// 6. 验证所有预期的 SQL 操作都已完成
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateMonitorTarget(t *testing.T) {
	// 1. 初始化Mock数据库
	mockDb, mock := tool.GetMysqlMock()

	// 替换全局DB实例（测试后恢复）
	originalDb := tool.Db
	tool.Db = mockDb
	defer func() { tool.Db = originalDb }()

	// 2. 准备测试数据
	input := MonitorTarget{
		Ip:         "192.168.1.1",
		Domain:     "",
		Condition:  "",
		IsDeleted:  0,
		CreateTime: time.Now(),
	}

	// 3. 设置Mock预期
	mock.ExpectBegin() // 预期事务开始（GORM默认启用事务）

	// 精确匹配INSERT语句
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `monitor_target` (`ip`,`domain`,`condition`,`is_deleted`,`create_time`) VALUES (?,?,?,?,?)")).
		WithArgs(input.Ip, input.Domain, input.Condition, input.IsDeleted, input.CreateTime).
		WillReturnResult(sqlmock.NewResult(1, 1)) // 模拟返回插入ID=1，影响1行

	mock.ExpectCommit() // 预期事务提交

	// 4. 执行创建操作
	result := tool.Db.Table("monitor_target").Create(&input)

	// 5. 验证结果
	assert.NoError(t, result.Error)
	assert.Equal(t, int64(1), result.RowsAffected)
	assert.Equal(t, 1, input.Id) // 验证自增ID被回填

	// 6. 验证所有预期SQL执行
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteMonitorTarget(t *testing.T) {
	// 1. 初始化Mock
	mockDb, mock := tool.GetMysqlMock()
	//defer mockDb.Close()

	tool.Db = mockDb // 替换全局实例

	// 2. 测试数据
	targetID := 1

	// 3. 设置Mock预期（完全匹配GORM生成的SQL格式）
	mock.ExpectBegin()

	// 注意：必须完全匹配GORM实际生成的SQL（包括所有反引号和空格）
	mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `monitor_target` SET `is_deleted`=? WHERE `id` = ?")). // 所有字段用反引号
		WithArgs(1, targetID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	// 4. 执行操作（使用一致的查询构造方式）
	result := tool.Db.Table("monitor_target").
		Where("`id` = ?", targetID). // 这里也要用反引号保持一致性
		Update("is_deleted", 1)

	// 5. 验证
	assert.NoError(t, result.Error)
	assert.Equal(t, int64(1), result.RowsAffected) // 关键断言
	assert.NoError(t, mock.ExpectationsWereMet())
}
