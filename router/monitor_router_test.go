package router

import (
	"MonitorProject/models"
	"MonitorProject/tool"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// 测试添加监控
func TestMonitorRouter(t *testing.T) {
	var configPath, configName, configType = "../config", "config", "yml"
	tool.ConfigInit(configPath, configName, configType)
	tool.InitDb()
	// 设置测试模式
	gin.SetMode(gin.TestMode)

	// 创建测试路由
	router := gin.Default()
	router.POST("/monitor", addMonitor)

	// 测试用例1: 正常添加监控
	t.Run("成功添加监控", func(t *testing.T) {
		monitor := models.MonitorTarget{
			Ip:         "192.168.302.1",
			CreateTime: time.Now(),
		}
		body, err := json.Marshal(monitor)
		if err != nil {
			t.Fatalf("序列化请求体失败: %v", err)
		}
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/monitor", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("创建请求失败: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		// 验证响应
		assert.Equal(t, http.StatusCreated, w.Code)
	})
	// 测试用例2: 重复添加
	t.Run("重复添加", func(t *testing.T) {
		monitor := models.MonitorTarget{
			Ip:         "1.0.0.1",
			CreateTime: time.Now(),
		}
		body, err := json.Marshal(monitor)
		if err != nil {
			t.Fatalf("序列化请求体失败: %v", err)
		}
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/monitor", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("创建请求失败: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		// 验证响应
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

// 测试查询monitor
func TestQueryHistoryDataByid(t *testing.T) {
	var configPath, configName, configType = "../config", "config", "yml"
	tool.ConfigInit(configPath, configName, configType)
	tool.InitDb()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/monitor/:id", getMonitor)
	// 测试用例1: 正常查询
	t.Run("成功查询历史数据", func(t *testing.T) {
		exceptedMonior := "{\"condition\":\"\",\"create_time\":\"2025-04-14T11:06:46+08:00\",\"domain\":\"\",\"id\":3,\"ip\":\"1.0.0.2\"}"

		// 查询历史数据
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/monitor/3", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		// 验证返回数据中没有changeCount字段
		assert.Equal(t, exceptedMonior, w.Body.String())
	})

	// 测试用例2: 查询不存在monitor
	t.Run("查询不存在monitor", func(t *testing.T) {
		// 查询历史数据
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/monitor/99", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		// 验证返回数据中没有changeCount字段
		assert.Contains(t, w.Body.String(), "Record not found")
	})
}

// 测试停止监控
func TestStopMonitor(t *testing.T) {
	var configPath, configName, configType = "../config", "config", "yml"
	tool.ConfigInit(configPath, configName, configType)
	tool.InitDb()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	//添加到监视
	router.DELETE("/monitor/:id", deleteMonitor)
	// 测试用例1: 正常停止监控
	t.Run("成功停止监控", func(t *testing.T) {
		// 停止监控
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/monitor/1", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	//测试用例2 停止不存在的监控
	t.Run("停止不存在监视器", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/monitor/99999", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Record not found")
	})

}

func TestQueryHistoryChangeDataByid(t *testing.T) {
	var configPath, configName, configType = "../config", "config", "yml"
	tool.ConfigInit(configPath, configName, configType)
	tool.InitDb()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/monitor/:id/history", getHistory)
	// 测试用例1: 正常查询
	t.Run("成功查询历史数据", func(t *testing.T) {
		// 查询历史数据
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/monitor/3/history", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "change_log")
	})

	// 测试用例2: 查询不存在的ID
	t.Run("查询未监控的ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/monitor/999/history", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Record not found")
	})
}
