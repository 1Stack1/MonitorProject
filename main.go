package main

import (
	"MonitorProject/jobs"
	"MonitorProject/router"
	"MonitorProject/tool"
	"fmt"
)

var configPath, configName, configType = "./config", "config", "yml"

func main() {
	err := tool.ConfigInit(configPath, configName, configType)
	if err != nil {
		fmt.Println(err)
		return
	}
	//初始化数据库
	err = tool.InitDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	//启动定时任务
	jobs.StartAssetUpdateJob()

	//暴露访问地址
	r := router.GetMonitorRouter()
	port := "8080"
	r.Run(":" + port)

}
