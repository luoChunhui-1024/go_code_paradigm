package packer

import (
	"encoding/json"
	"github.com/go_code_paradigm/moderate_regular/biz/entity"
	"github.com/go_code_paradigm/moderate_regular/biz/model"
)

// PackOrder 包装单个订单
func PackOrder(order entity.ProductOrder) model.OrderInfo {
	var extensions entity.OrderExtensions
	_ = json.Unmarshal([]byte(order.Extensions), &extensions)

	return model.OrderInfo{
		ID:        order.ID,
		OrderID:   order.OrderID,
		Customer:  order.CustomerID,
		Product:   order.ProductID,
		Amount:    order.Amount,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		Note:      extensions.Note,
	}
}

// PackOrders 包装多个订单
func PackOrders(orders []entity.ProductOrder) []model.OrderInfo {
	var responses []model.OrderInfo
	for _, order := range orders {
		responses = append(responses, PackOrder(order))
	}
	return responses
}
