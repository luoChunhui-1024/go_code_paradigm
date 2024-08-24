package entity

import "time"

const ProductOrderTableName = "product_order"

// ProductOrder 代表一个订单的数据结构
type ProductOrder struct {
	ID         int       `json:"id" gorm:"column:id"`
	CustomerID int       `json:"customer_id" gorm:"column:customer_id"`
	Amount     float64   `json:"amount" gorm:"column:amount"`
	Status     string    `json:"status" gorm:"column:status"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
}

func (p *ProductOrder) TableName() string {
	return ProductOrderTableName
}
