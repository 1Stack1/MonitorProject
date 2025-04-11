package models

import (
	"MonitorProject/tool"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
	"time"
)

func TestGetHistory_Success(t *testing.T) {
	mockDb, mock := tool.GetMysqlMock()

	// 替换全局的 tool.Db 为 mock 实例
	originalDb := tool.Db
	tool.Db = mockDb
	defer func() { tool.Db = originalDb }() // 测试结束后恢复原 DB

	time, err := time.Parse("2006-01-02 15:04:05", "2025-04-11 10:19:14")
	if err != nil {
		log.Fatal(err)
	}
	// 2. 定义预期数据
	expectedHistory := MonitorHistory{
		Id:               1,
		TargetId:         2,
		MonitorStartTime: time,
		AssetCount:       3524,
		ChangedCount:     0,
		IsDeleted:        0,
		CreateTime:       time,
	}

	// 3. 设置 Mock 预期, 预期执行的 SQL 查询
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `monitor_history` WHERE id = ? and is_deleted = ?")).
		// 匹配查询参数
		WithArgs(1, 0).
		// 返回模拟数据行
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "target_id", "monitor_start_time", "asset_count", "changed_count", "is_deleted", "create_time"}).
				AddRow(expectedHistory.Id, expectedHistory.TargetId, expectedHistory.MonitorStartTime,
					expectedHistory.AssetCount, expectedHistory.ChangedCount, expectedHistory.IsDeleted, expectedHistory.CreateTime),
		)

	// 4. 调用被测方法
	var history MonitorHistory
	result := tool.Db.Debug().
		Table("monitor_history").
		Where("id = ? and is_deleted = ?", 1, 0).
		Find(&history)

	// 5. 断言结果
	assert.NoError(t, result.Error)                 // 验证没有错误
	assert.Equal(t, expectedHistory.Id, history.Id) // 验证返回的 ID 符合预期
	assert.Equal(t, expectedHistory.TargetId, history.TargetId)
	assert.Equal(t, expectedHistory.MonitorStartTime, history.MonitorStartTime)
	assert.Equal(t, expectedHistory.AssetCount, history.AssetCount)
	assert.Equal(t, expectedHistory.ChangedCount, history.ChangedCount)
	assert.Equal(t, expectedHistory.IsDeleted, history.IsDeleted)
	assert.Equal(t, expectedHistory.CreateTime, history.CreateTime)

	// 6. 验证所有预期的 SQL 操作都已完成
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetLastHistory_Success(t *testing.T) {
	// 1. 初始化Mock数据库
	mockDb, mock := tool.GetMysqlMock()
	//defer mockDb.Close()

	// 替换全局DB实例（测试后恢复）
	originalDb := tool.Db
	tool.Db = mockDb
	defer func() { tool.Db = originalDb }()

	// 2. 准备测试数据
	testTargetID := 1
	expectedCount := 3524

	// 3. 设置Mock预期
	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT asset_count FROM monitor_history WHERE target_id = ? ORDER BY id desc LIMIT 1")).
		WithArgs(testTargetID). // 验证参数是否正确
		WillReturnRows(
			sqlmock.NewRows([]string{"asset_count"}).
				AddRow(expectedCount), // 模拟返回单行单列数据
		)

	// 4. 执行查询
	var lastHistoryCount int
	err := tool.Db.Debug().
		Raw("SELECT asset_count FROM monitor_history WHERE target_id = ? ORDER BY id desc LIMIT 1", testTargetID).
		Scan(&lastHistoryCount).Error

	// 5. 验证结果
	assert.NoError(t, err)
	assert.Equal(t, expectedCount, lastHistoryCount)
	assert.NoError(t, mock.ExpectationsWereMet())
}
