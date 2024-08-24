package service

import (
	"github.com/go_code_paradigm/moderate_regular/biz/dal"
	"github.com/go_code_paradigm/moderate_regular/biz/model"
	"github.com/go_code_paradigm/moderate_regular/biz/packer"
)

// OrderService 订单服务
type OrderService struct{}

// NewOrderService 创建新的订单服务
func NewOrderService() *OrderService {
	return &OrderService{}
}

// GetOrders 条件查询订单
func (s *OrderService) GetOrders(customerIDs []int, orderIDs []int, minAmount, maxAmount float64, startDate, endDate string) ([]model.OrderInfo, error) {
	orders, err := dal.GetOrders(customerIDs, orderIDs, minAmount, maxAmount, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return packer.PackOrders(orders), nil
}

// UpdateOrder 更新订单状态
func (s *OrderService) UpdateOrder(id int, status, extensions string) (model.OrderInfo, error) {
	order, err := dal.UpdateOrder(id, status, extensions)
	if err != nil {
		return model.OrderInfo{}, err
	}
	return packer.PackOrder(order), nil
}
