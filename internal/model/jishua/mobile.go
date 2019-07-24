package jishua

import (
	"github.com/hjd919/server/internal/model"
)

// Mobile
type Mobile struct {
	model.MID
	UniqueNo       string `bson:"unique_no"`        // 手机标识
	Group          int    `bson:"group"`            // 手机组号
	UseEndTime     int    `bson:"use_end_time"`     // 使用结束时间、大于当前时间为未使用
	LastAccessTime int    `bson:"last_access_time"` // 最后访问时间，小于5分钟为有效
}
