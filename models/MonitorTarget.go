package models

type MonitorTarget struct {
	Id        int
	Ip        string
	Domain    string
	Condition string
	IsDeleted int
}
