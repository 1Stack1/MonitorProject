package main

import (
	"MonitorProject/router"
	"MonitorProject/tool"
)

var configPath, configName, configType = "./config", "config", "yml"

func main() {
	tool.ConfigInit(configPath, configName, configType)
	//tool.SendMail()
	/*//启动定时任务
	jobs.StartAssetUpdateJob()*/

	//初始化数据库
	tool.InitDb()

	//暴露访问地址
	r := router.GetMonitorRouter()
	port := "8080"
	r.Run(":" + port)

}
