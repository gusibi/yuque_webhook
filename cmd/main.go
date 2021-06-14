package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gusibi/yuque_webhook/api"
	"time"
)

func main() {
	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	r.GET("/ping", api.Ping)
	r.POST("/api/lark_webhook/:hook_id", api.LarkWebHook)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
