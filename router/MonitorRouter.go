package router

import (
	"MonitorProject/models"
	"MonitorProject/tool"
	"github.com/gin-gonic/gin"
	"net/http"
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
	result := tool.Db.Table("monitor_target").Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": input.Id})
	//_ := models.MonitorTarget{}
}

// 获取监控目标详情
func getMonitor(c *gin.Context) {

}

// 删除监控目标
func deleteMonitor(c *gin.Context) {

}

// 获取资产变化历史
func getHistory(c *gin.Context) {

}
