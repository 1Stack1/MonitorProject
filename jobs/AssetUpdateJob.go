package jobs

import (
	"MonitorProject/models"
	"MonitorProject/tool"
	"github.com/robfig/cron/v3"
	"log"
)

/*
定时任务：每天更新资产数据
*/
func StartAssetUpdateJob() {
	c := cron.New(cron.WithSeconds()) // 启用秒级解析
	_, err := c.AddFunc("0 0 * * *", func() {
		updateAssets()
	})

	if err != nil {
		log.Fatal("Error scheduling job:", err)
	}

	c.Start()

}

func updateAssets() {
	var input models.MonitorTarget
	input = models.MonitorTarget{Ip: "1.1.1.1"}

	_ = tool.Db.Table("monitor_target").Create(&input)

}
