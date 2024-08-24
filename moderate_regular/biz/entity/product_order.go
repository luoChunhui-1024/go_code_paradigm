package entity

import "time"

const ProductOrderTableName = "product_order"

// ProductOrder 订单表 model
type ProductOrder struct {
	ID         int       `json:"id" gorm:"column:id"`
	OrderID    int       `json:"order_id" gorm:"column:order_id"`
	CustomerID int       `json:"customer_id" gorm:"column:customer_id"`
	ProductID  int       `json:"product_id" gorm:"column:product_id"`
	Amount     float64   `json:"amount" gorm:"column:amount"`
	Status     string    `json:"status" gorm:"column:status"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
	Extensions string    `json:"extensions" gorm:"column:extensions"` // 存储 JSON 格式的扩展字段，如备注和快照
}

// OrderExtensions 订单的扩展字段
type OrderExtensions struct {
	Note     string `json:"note"`
	Snapshot string `json:"snapshot"`
}

func (p *ProductOrder) TableName() string {
	return ProductOrderTableName
}
