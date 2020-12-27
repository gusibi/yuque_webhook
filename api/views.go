package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gusibi/yuque_webhook/internal/config"
	"github.com/gusibi/yuque_webhook/internal/yuque"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func WebHook(c *gin.Context) {

	var req yuque.WebHookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.Printf("req:%+v \n", req)
	if err := yuque.RequestValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if v, err := json.Marshal(req); err == nil {
		fmt.Printf("%s\n", v)
	}
	hookId := config.Settings.LarkHookId
	larkNotify := &yuque.LarkWebHook{
		MessageType:    yuque.PostMessage,
		HookId:         hookId,
		DefaultTimeout: 2,
	}
	webhook := yuque.NewWebHook()
	webhook.Register(c, larkNotify)
	webhook.Notify(c, &req)
	c.JSON(http.StatusOK, req)
}
