package main

import (
	"MonitorProject/router"
	"MonitorProject/tool"
)

func main() {
	tool.InitDb()

	r := router.GetMonitorRouter()
	port := "8080"
	r.Run(":" + port)
}
