package api

import "github.com/globalsign/mgo/bson"

type Agent struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Hostname string        `json:"hostname" bson:"hostname"`
	IP       string        `json:"ip" bson:"ip"`
	Ct       int64         `json:"ct" bson:"ct"`
}
