package jobs

import (
	"MonitorProject/tool"
	"fmt"
	"testing"
)

func TestAssetMonitor_Success(t *testing.T) {
	var configPath, configName, configType = "../config", "config", "yml"
	err := tool.ConfigInit(configPath, configName, configType)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tool.InitDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ChangedThresholdInit()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = UrlInit()
	if err != nil {
		fmt.Println(err)
		return
	}
	AssetMoniter()
}
