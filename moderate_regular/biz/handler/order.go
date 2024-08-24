package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go_code_paradigm/moderate_regular/biz/service"
	"net/http"
	"strconv"
)

// GetOrdersHandler 处理获取订单的请求
func GetOrdersHandler(c *gin.Context) {
	status := c.Query("status")
	svr := service.NewOrderService()
	orders, err := svr.GetOrders(status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// UpdateOrderHandler 处理更新订单的请求
func UpdateOrderHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order ID"})
		return
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	svr := service.NewOrderService()
	order, err := svr.UpdateOrder(id, req.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}
