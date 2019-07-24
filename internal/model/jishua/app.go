package jishua

import "github.com/hjd919/server/internal/model"

// App
type App struct {
	model.MID
	BundleID        string `bson:"bundleId"`
	Channel         string `bson:"channel"`
	Appid           string `bson:"appid"`
	AppName         string `bson:"app_name"`
	CallbackTime    string `bson:"callback_time"`
	MaxCallbackTime string `bson:"max_callback_time"`
	OpenTime        string `bson:"open_time"`
	MaxOpenTime     string `bson:"max_open_time"`
	process         string `bson:"process"`
	taskType        string `bson:"taskType"`
	needClean       string `bson:"needClean"`
	VersionLimit    string `bson:"version_limit"`
	DoubleOpen      string `bson:"double_open"`
	CallbackURL     string `bson:"callback_url"`
	ActiveURL       string `bson:"active_url"`
	QueryURL        string `bson:"query_url"`
	ClickURL        string `bson:"click_url"`
}
