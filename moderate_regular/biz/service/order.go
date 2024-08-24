package service

import (
	"github.com/go_code_paradigm/moderate_regular/biz/dal"
	"github.com/go_code_paradigm/moderate_regular/biz/entity"
)

// OrderService 订单服务
type OrderService struct{}

// NewOrderService 创建新的订单服务
func NewOrderService() *OrderService {
	return &OrderService{}
}

// GetOrders 条件查询订单
func (s *OrderService) GetOrders(status string) ([]entity.ProductOrder, error) {
	return dal.GetOrders(status)
}

// UpdateOrder 更新订单状态
func (s *OrderService) UpdateOrder(id int, status string) (entity.ProductOrder, error) {
	return dal.UpdateOrder(id, status)
}
