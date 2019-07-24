package jishua

import "github.com/hjd919/server/internal/model"

type KeywordApp struct {
	model.MID
	BundleID string `bson:"bundleId"`
	Channel  string `bson:"channel"`
	Appid    string `bson:"appid"`
	AppName  string `bson:"app_name"`
}

// Keywords
type Keyword struct {
	model.MID
	keyword string     `bson:"keyword"`
	App     KeywordApp `bson:"app"` // 手机标识
}
