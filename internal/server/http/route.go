package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func initRouter(e *bm.Engine) {
	e.POST("/record", svc.Record)
	// student := e.Group("/api/v1")
	// {
	// 	//获取学生列表
	// 	student.GET("/test", ListStudent)
	// }
}
