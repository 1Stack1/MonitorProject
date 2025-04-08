package models

type MonitorHistory struct {
	Id               int
	TargetId         int
	MonitorStartTime string
	AssetCount       string
	ChangedLog       string
	IsDeleted        int
	CreatTime        string
}
