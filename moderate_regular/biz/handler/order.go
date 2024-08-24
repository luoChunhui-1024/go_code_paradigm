package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go_code_paradigm/moderate_regular/biz/model"
	"github.com/go_code_paradigm/moderate_regular/biz/service"
	"net/http"
)

// GetOrdersHandler 处理获取订单的请求
func GetOrdersHandler(c *gin.Context) {
	svr := service.NewOrderService()
	var params model.OrderQueryReq

	// 绑定并验证请求参数
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 调用服务层获取订单数据
	orders, err := svr.GetOrders(params.CustomerIDs, params.OrderIDs, params.MinAmount, params.MaxAmount, params.StartDate, params.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

// UpdateOrderHandler 处理更新订单的请求
func UpdateOrderHandler(c *gin.Context) {
	var req model.UpdateOrderreq
	// 绑定并验证请求参数
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	svr := service.NewOrderService()
	order, err := svr.UpdateOrder(req.OrderID, req.Status, req.Extensions)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}
