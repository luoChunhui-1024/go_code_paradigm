package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go_code_paradigm/moderate_regular/biz/handler"
	"net/http"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func register(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	{
		// 订单相关路由
		r.GET("/orders", handler.GetOrdersHandler)
		r.PUT("/orders/:id", handler.UpdateOrderHandler)
	}
	{
		// 商品相关路由

	}
	// 绑定JSON的例子 ({"user": "manu", "password": "123"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var req Login
		if c.BindJSON(&req) == nil {
			if req.User == "manu" && req.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	// 绑定普通表单的例子 (user=manu&password=123)
	r.POST("/loginForm", func(c *gin.Context) {
		var formReq Login
		// 根据请求头中 content-type 自动推断.
		if c.Bind(&formReq) == nil {
			if formReq.User == "manu" && formReq.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})
}
