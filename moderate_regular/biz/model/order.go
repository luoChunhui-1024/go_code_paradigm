package model

import "time"

// OrderQueryReq 是订单查询的请求参数结构体
type OrderQueryReq struct {
	CustomerIDs []int   `form:"customer_ids" binding:"omitempty,dive,gt=0"`
	OrderIDs    []int   `form:"order_ids" binding:"omitempty,dive,gt=0"`
	MinAmount   float64 `form:"min_amount" binding:"omitempty,gte=0"`
	MaxAmount   float64 `form:"max_amount" binding:"omitempty,gte=0"`
	StartDate   string  `form:"start_date" binding:"omitempty,datetime=2006-01-02"`
	EndDate     string  `form:"end_date" binding:"omitempty,datetime=2006-01-02"`
}

// OrderInfo 是返回给前端的订单结构体
type OrderInfo struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"order_id"`
	Customer  int       `json:"customer"`
	Product   int       `json:"product"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Note      string    `json:"note,omitempty"`
	Snapshot  string    `json:"snapshot,omitempty"`
}

type UpdateOrderreq struct {
	OrderID    int    `json:"order_id"`
	Status     string `json:"status"`
	Extensions string `json:"extensions"`
}
