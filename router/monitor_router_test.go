package router

import (
	"MonitorProject/models"
	"MonitorProject/tool"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddMonitor_Unit(t *testing.T) {
	// 1. 创建Mock DB
	mockDB, mock := tool.GetMysqlMock()

	// 2. 设置Mock预期
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `monitor_target`.*").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// 3. 创建测试请求
	payload := models.MonitorTarget{
		Ip:     "192.168.1.1",
		Domain: "example.com",
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/monitor", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// 4. 创建路由并注入Mock DB
	router := gin.Default()
	router.POST("/monitor", func(c *gin.Context) {
		tool.Db = mockDB // 注入Mock
		addMonitor(c)
	})

	// 5. 发起请求
	router.ServeHTTP(w, req)

	// 6. 验证
	assert.Equal(t, http.StatusCreated, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(1), response["id"]) // JSON数字会被解码为float64
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetMonitor(t *testing.T) {
	// 初始化Mock
	mockDB, mock := tool.GetMysqlMock()
	tool.Db = mockDB // 替换全局实例

	// 测试用例：成功查询
	t.Run("Success", func(t *testing.T) {
		// 设置Mock预期
		mock.ExpectQuery(regexp.QuoteMeta(
			"SELECT * FROM `monitor_target` WHERE id = ? AND is_deleted = ?")).
			WithArgs("1", 0).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "ip", "domain"}).
					AddRow(1, "192.168.1.1", "example.com"),
			)

		// 创建测试请求
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/monitor/1", nil)
		router := gin.Default()
		router.GET("/monitor/:id", getMonitor)
		router.ServeHTTP(w, req)

		// 验证
		assert.Equal(t, http.StatusOK, w.Code)
		expected := `{
            "id": 1,
            "ip": "192.168.1.1",
            "domain": "example.com",
            "condition": "",
            "create_time": "0001-01-01T00:00:00Z"
        }`
		assert.JSONEq(t, expected, w.Body.String())
	})

	// 测试用例：记录不存在
	t.Run("Not Found", func(t *testing.T) {
		mock.ExpectQuery("SELECT.*").
			WillReturnRows(sqlmock.NewRows([]string{"id"})) // 返回空结果

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/monitor/999", nil)
		GetMonitorRouter().ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Record not found")
	})
}

func TestDeleteMonitor(t *testing.T) {
	mockDB, mock := tool.GetMysqlMock()

	tool.Db = mockDB

	// 成功删除用例
	t.Run("Success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(
			"UPDATE `monitor_target` SET `is_deleted` = ? WHERE `id` = ?")).
			WithArgs(1, "1").
			WillReturnResult(sqlmock.NewResult(0, 1)) // 影响1行
		mock.ExpectCommit()

		w := httptest.NewRecorder()
		req := httptest.NewRequest("DELETE", "/monitor/1", nil)
		router := gin.Default()
		router.DELETE("/monitor/:id", deleteMonitor)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Monitor stopped")
	})

	// 记录不存在用例
	t.Run("Not Found", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE.*").
			WillReturnResult(sqlmock.NewResult(0, 0)) // 影响0行
		mock.ExpectRollback()

		w := httptest.NewRecorder()
		req := httptest.NewRequest("DELETE", "/monitor/999", nil)
		GetMonitorRouter().ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "this have changed")
	})
}

func TestGetHistory(t *testing.T) {
	mockDB, mock := tool.GetMysqlMock()

	tool.Db = mockDB

	// 成功查询用例
	t.Run("Success", func(t *testing.T) {
		// 准备模拟数据
		mockTime := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)
		mock.ExpectQuery(regexp.QuoteMeta(
			"SELECT * FROM `monitor_history` WHERE target_id = ? AND is_deleted = ?")).
			WithArgs("1", 0).
			WillReturnRows(
				sqlmock.NewRows([]string{"target_id", "changed_count", "monitor_start_time"}).
					AddRow(1, 5, mockTime).
					AddRow(1, 3, mockTime.AddDate(0, 0, -1)),
			)

		// 发送请求
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/monitor/1/history", nil)
		router := gin.Default()
		router.GET("/monitor/:id/history", getHistory)
		router.ServeHTTP(w, req)

		// 验证响应
		assert.Equal(t, http.StatusOK, w.Code)
		expected := `{
            "id": 1,
            "change_log": [
                {"monitor_date": "2023:6:15", "changed_count": 5},
                {"monitor_date": "2023:6:14", "changed_count": 3}
            ]
        }`
		assert.JSONEq(t, expected, w.Body.String())
	})

	// 空记录用例
	t.Run("Empty History", func(t *testing.T) {
		mock.ExpectQuery("SELECT.*").
			WillReturnRows(sqlmock.NewRows([]string{"target_id"}))

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/monitor/1/history", nil)
		GetMonitorRouter().ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
