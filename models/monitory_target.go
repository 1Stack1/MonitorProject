package models

import (
	"time"
)

type MonitorTarget struct {
	Id         int
	Ip         string
	Domain     string
	Condition  string
	IsDeleted  int
	CreateTime time.Time
}
