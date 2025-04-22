package tool

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Number string
}

var (
	lastUpdate   time.Time
	ignoredFiles = []string{".tmp", "~"} // 需要忽略的文件后缀
)

var viperContext *viper.Viper

/*
初始化viper
*/
func ConfigInit(configPath string, configName string, configType string) error {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		// 忽略临时文件
		for _, suffix := range ignoredFiles {
			if strings.HasSuffix(e.Name, suffix) {
				return
			}
		}

		// 防抖：500ms 内仅处理一次
		if time.Since(lastUpdate) < 500*time.Millisecond {
			return
		}
		lastUpdate = time.Now()

		fmt.Println("config file changed:", e.Name)
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
	})
	viperContext = v
	return nil
}

/*
读取数据库连接
*/
func ConfigReadDatabaseConnection() (string, error) {
	databaseConnection := viperContext.GetString("database_connection")
	if databaseConnection == "" {
		return "", fmt.Errorf("数据库连接必须定义")
	}
	return databaseConnection, nil
}

/*
读取邮箱密码
*/
func ConfigReadEmailPassword() (string, error) {
	emailPassword := viperContext.GetString("email_password")
	if emailPassword == "" {
		return "", fmt.Errorf("邮箱密码必须定义")
	}
	return emailPassword, nil
}

func ConfigReadFromEmailUser() (string, error) {
	FromEmailUser := viperContext.GetString("from_email_user")
	if FromEmailUser == "" {
		return "", fmt.Errorf("发件人必须定义")
	}
	return FromEmailUser, nil
}

func ConfigReadToEmailUser() (string, error) {
	ToEmailUser := viperContext.GetString("to_email_user")
	if ToEmailUser == "" {
		return "", fmt.Errorf("送件人必须定义")
	}
	return ToEmailUser, nil
}

/*
读取用户key
*/
func ConfigReadUserKey() (string, error) {
	userKey := viperContext.GetString("user_key")
	if userKey == "" {
		return "", fmt.Errorf("userkey必须定义")
	}
	return userKey, nil
}

/*
读取阈值
*/
func ConfigReadChangedThreshold() (int, error) {
	changedThresholdStr := viperContext.GetString("changed_threshold")
	if changedThresholdStr == "" {
		return 10, nil
	}
	changedThreshold, err := strconv.Atoi(changedThresholdStr)
	if err != nil {
		return 0, fmt.Errorf("changedThreshold转换为整数错误: %w", err)
	}
	return changedThreshold, nil
}
