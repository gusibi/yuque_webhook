package main

// https://github.com/go-swagger/go-swagger/issues/962

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"

	"github.com/gusibi/yuque_webhook/api"

	"github.com/aws/aws-lambda-go/events"
	scf "github.com/tencentyun/scf-go-lib/cloudevents/scf"
	"github.com/tencentyun/scf-go-lib/cloudfunction"

	"github.com/gusibi/yuque_webhook/cmd/scf/httpadapter"
)

var httpAdapter *httpadapter.GinLambda

func init() {
	log.Println("start server...")
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
	r.POST("/api/webhook", api.WebHook)
	httpAdapter = httpadapter.New(r)
	log.Println("adapter: ", httpAdapter)
}

// Handler go swagger aws lambda handler
func Handler(req events.APIGatewayProxyRequest) (scf.APIGatewayProxyResponse, error) {

	return httpAdapter.Proxy(req)
}

func main() {
	cloudfunction.Start(Handler)
}
