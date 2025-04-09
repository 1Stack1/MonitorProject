package main

/*import (
	"MonitorProject/models"
	"fmt"
	"net/http"
"net/http/httptest"
"testing"
"time"

"github.com/gin-gonic/gin"
"github.com/stretchr/testify/assert"
"gorm.io/gorm"

)

func setupTestDB() *gorm.DB {

}

func TestAddMonitor(t *testing.T) {

}

func TestGetHistory(t *testing.T) {
	db = setupTestDB()

	// 创建测试数据
	target := models.MonitorTarget{Target: "test.com", IsActive: true}
	db.Create(&target)

	history := models.AssetHistory{
		TargetID:   target.ID,
		CheckDate:  time.Now().UTC().Truncate(24 * time.Hour),
		AssetCount: 100,
	}
	db.Create(&history)

	// 创建路由
	r := gin.Default()
	r.GET("/monitor/:id/history", getHistory)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/monitor/%d/history", target.ID), nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"assetCount":100`)
}
*/
