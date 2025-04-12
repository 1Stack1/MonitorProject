package jobs

import (
	"MonitorProject/models"
	"MonitorProject/tool"
	"encoding/base64"
	"fmt"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var url string

/*
定时任务：每天更新资产数据
*/
func StartAssetUpdateJob() {
	c := cron.New(cron.WithSeconds()) // 启用秒级解析
	_, err := c.AddFunc("0 0 * * *", func() {

	})

	if err != nil {
		log.Fatal("Error scheduling job:", err)
	}

	c.Start()

}

func UrlInit() error {
	urlPrefix := "https://fofa.info/api/v1/search/all?email=&full=true&page=1&size=1&key="
	key, err := tool.ConfigReadUserKey()
	if err != nil {
		return err
	}
	url = urlPrefix + key
	return nil
}

/*
资产监控
*/
func AssetMoniter() {
	//todo 添加个测试go文件
	var targets []models.MonitorTarget // 改为切片类型

	if result := tool.Db.Debug().Table("monitor_target").Where("is_deleted = ?", 0).Find(&targets); result.Error != nil {
		return
	}

	for _, target := range targets {
		//fofa调用
		fullUrl := buildFullUrl(target)
		size, err := fofa(fullUrl)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//计算资产数量变化
		var lastHistoryCount int
		if result := tool.Db.Debug().
			Raw("SELECT asset_count FROM monitor_history WHERE target_id = ? ORDER BY id desc LIMIT 1", target.Id).
			Scan(&lastHistoryCount); result.Error != nil {
			continue
		}
		//todo 计算第一次监控
		//保存记录
		history := models.MonitorHistory{
			TargetId:         target.Id,
			AssetCount:       size,
			MonitorStartTime: time.Now(),
			ChangedCount:     size - lastHistoryCount,
			CreateTime:       time.Now(),
		}
		if result := tool.Db.Debug().Table("monitor_history").Create(&history); result.Error != nil {
			continue
		}
		//todo 如果changedCount超过阈值发送邮件
	}
}

/*
建立完整url
*/
func buildFullUrl(target models.MonitorTarget) string {
	//todo 将url用&拼起来
	if target.Ip != "" {
		queryContent := "ip=\"" + target.Ip + "\""
		encoded := base64QueryArg(queryContent)
		return concatenatedQueryCondition(encoded)
	} else if target.Domain != "" {
		queryContent := "domain=\"" + target.Domain + "\""
		encoded := base64QueryArg(queryContent)
		return concatenatedQueryCondition(encoded)
	} else {
		queryContent := target.Condition
		encoded := base64QueryArg(queryContent)
		return concatenatedQueryCondition(encoded)
	}
}

func base64QueryArg(QueryContent string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(QueryContent))
	encoded = "&qbase64=" + encoded
	return encoded
}

/*
拼接查询条件
*/
func concatenatedQueryCondition(condition string) string {
	resultUrl := url + condition
	return resultUrl
}

func fofa(url string) (int, error) {
	var client *http.Client
	client = &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fofaRes, err := tool.FofaResJsonDes(string(body))
	if err != nil {
		return 0, err
	}
	//fmt.Printf("内容: %s\n", string(body))
	return fofaRes.Size, nil
}
