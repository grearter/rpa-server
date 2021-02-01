package api

import "github.com/globalsign/mgo/bson"

type RobotStatus string

const (
	RobotStatusStopped RobotStatus = "stopped"
	RobotStatusRunning RobotStatus = "running"
)

type Robot struct {
	ID       bson.ObjectId   `json:"id" bson:"_id"`
	AgentID  bson.ObjectId   `json:"agentId" bson:"agentId"`
	Filepath string          `json:"filepath" bson:"filepath"`
	Status   RobotStatus     `json:"status"`
	Ct       int64           `json:"ct" bson:"ct"`
	TaskIDs  []bson.ObjectId `json:"taskIds" bson:"taskIds"`
}

type RobotTask struct {
	ID         bson.ObjectId   `json:"id" bson:"_id"`
	RobotID    bson.ObjectId   `json:"robotId" bson:"robotId"`
	CreateID   string          `json:"createId" bson:"createId"`
	CreateName string          `json:"createName" bson:"createName"`
	Ct         int64           `json:"ct" bson:"ct"`
	Ut         int64           `json:"ut" bson:"ut"`
	Messages   []*RobotMessage `json:"messages" bson:"messages"`
}

type RobotMessage struct {
	Process string `json:"process" bson:"process"` // 机器人流程名称
	Level   string `json:"level" bson:"level"`     // 日志级别
	Ct      int64  `json:"ct" bson:"ct"`           // 日志时间
	Content string `json:"content" bson:"content"` // 日志描述
}
