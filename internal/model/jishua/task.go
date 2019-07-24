package jishua

import "github.com/hjd919/server/internal/model"

type TaskMobile struct {
	model.MID
	UniqueNo string `bson:"unique_no"` // 手机标识
}

// Task
type Task struct {
	model.MID
	App        App
	Mobile     TaskMobile
	StartTime  int `bson:"start_time"`
	EndTime    int `bson:"end_time"`
	IsFinish   int `bson:"is_finish"`
	FetchNum   int `bson:"fetch_num"`
	OrderNum   int `bson:"order_num"`
	ReturnNum  int `bson:"return_num"`
	SuccessNum int `bson:"success_num"`
}
