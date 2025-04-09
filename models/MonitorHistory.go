package models

import "time"

type MonitorHistory struct {
	Id               int
	TargetId         int
	MonitorStartTime string
	AssetCount       int
	ChangedCount     int
	ChangedLog       string
	IsDeleted        int
	CreatTime        time.Time
}
