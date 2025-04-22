package router

import (
	"MonitorProject/models"
	"MonitorProject/models/dto"
	"MonitorProject/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetMonitorRouter() *gin.Engine {
	r := gin.Default()

	// API路由
	r.POST("/monitor", addMonitor)
	r.GET("/monitor/:id", getMonitor)
	r.DELETE("/monitor/:id", deleteMonitor)
	r.GET("/monitor/:id/history", getHistory)

	return r
}

// 添加监控目标
func addMonitor(c *gin.Context) {
	var input models.MonitorTarget
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.CreateTime = time.Now()
	result := tool.Db.Table("monitor_target").Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": input.Id})
}

// 获取监控目标详情
func getMonitor(c *gin.Context) {
	var target models.MonitorTarget
	id := c.Param("id")

	if result := tool.Db.Debug().Table("monitor_target").Where("id = ? and is_deleted = ?", id, 0).Find(&target); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if target.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          target.Id,
		"ip":          target.Ip,
		"domain":      target.Domain,
		"condition":   target.Condition,
		"create_time": target.CreateTime,
	})
}

// 删除监控目标
func deleteMonitor(c *gin.Context) {
	id := c.Param("id")
	result := tool.Db.Table("monitor_target").Model(&models.MonitorTarget{}).Where("id = ?", id).Update("is_deleted", 1)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "this have changed/Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Monitor stopped"})
}

// 获取资产变化历史
func getHistory(c *gin.Context) {
	var histories []models.MonitorHistory
	id := c.Param("id")

	if result := tool.Db.Debug().Table("monitor_history").
		Where("target_id = ? and is_deleted = ?", id, 0).Find(&histories); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if len(histories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	var assetChangeLog []dto.AssetChangeLog
	for _, history := range histories {
		assetChangeLog = append(assetChangeLog, dto.AssetChangeLog{
			MonitorDate:  fmt.Sprintf("%d:%d:%d", history.MonitorStartTime.Year(), history.MonitorStartTime.Month(), history.MonitorStartTime.Day()),
			ChangedCount: history.ChangedCount,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         histories[0].TargetId,
		"change_log": assetChangeLog,
	})
}
