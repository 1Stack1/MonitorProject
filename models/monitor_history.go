package models

import "time"

type MonitorHistory struct {
	Id               int
	TargetId         int
	MonitorStartTime time.Time
	AssetCount       int
	ChangedCount     int
	ChangedLog       string
	IsDeleted        int
	CreateTime       time.Time
}
