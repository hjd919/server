package http

import (
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

// func initRouter(e *gin.Engine) {
// 	// e.GET("/swagger/*any", ginSwagge r.WrapHandler(swaggerFiles.Handler))
// 	student := e.Group("/api/v1")
// 	{
// 		//获取学生列表
// 		student.GET("/test", route.ListStudent)
// 	}
// }

func initRouter(e *bm.Engine) {
	// e.Ping(ping)
	student := e.Group("/api/v1")
	{
		//获取学生列表
		student.GET("/test", ListStudent)
	}
}
