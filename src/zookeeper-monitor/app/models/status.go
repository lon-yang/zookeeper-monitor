package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Status :Zookeeper Server Status
type Status struct {
	ID                      int       `orm:"column(ID)"`
	ServerID                int       `orm:"column(ServerID)"`
	Version                 string    `orm:"column(Version);"`
	AvgLatency              int64     `orm:"column(AvgLatency);"`
	MaxLatency              int64     `orm:"column(MaxLatency);"`
	MinLatency              int64     `orm:"column(MinLatency);"`
	PacketsReceived         int64     `orm:"column(PacketsReceived);"`
	PacketsSend             int64     `orm:"column(PacketsSend);"`
	NumAliveConnections     int64     `orm:"column(NumAliveConnections);"`
	OutstandingRequests     int64     `orm:"column(OutstandingRequests);"`
	ServerState             string    `orm:"column(ServerState);"`
	ZnodeCount              int64     `orm:"column(ZnodeCount);"`
	WatchCount              int       `orm:"column(WatchCount);"`
	EphemeralsCount         int       `orm:"column(EphemeralsCount);"`
	ApproximateDataSize     int       `orm:"column(ApproximateDataSize);"`
	OpenFileDescriptorCount int       `orm:"column(OpenFileDescriptorCount);"`
	MaxFileDescriptorCount  int       `orm:"column(MaxFileDescriptorCount);"`
	Followers               int       `orm:"column(Followers);"`
	SyncedFollowers         int       `orm:"column(SyncedFollowers);"`
	PendingSyncs            int       `orm:"column(PendingSyncs);"`
	InDate                  time.Time `orm:"column(InDate);type(datetime)"`
}

//AddStatus : Add Status
func AddStatus(s *Status) (int64, error) {
	s.InDate = time.Now()
	return orm.NewOrm().Insert(s)
}