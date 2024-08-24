package dal

import (
	"database/sql"
	"errors"
	"github.com/go_code_paradigm/moderate_regular/biz/entity"
	"github.com/go_code_paradigm/moderate_regular/pkg/db"
)

// GetOrders 条件查询订单
func GetOrders(status string) ([]entity.ProductOrder, error) {
	var rows *sql.Rows
	var err error

	if status == "" {
		rows, err = db.DB.Query("SELECT id, customer_id, amount, status, created_at FROM orders")
	} else {
		rows, err = db.DB.Query("SELECT id, customer_id, amount, status, created_at FROM orders WHERE status = ?", status)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.ProductOrder
	for rows.Next() {
		var order entity.ProductOrder
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.Amount, &order.Status, &order.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// UpdateOrder 更新订单状态
func UpdateOrder(id int, status string) (entity.ProductOrder, error) {
	result, err := db.DB.Exec("UPDATE orders SET status = ? WHERE id = ?", status, id)
	if err != nil {
		return entity.ProductOrder{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return entity.ProductOrder{}, err
	}

	if rowsAffected == 0 {
		return entity.ProductOrder{}, errors.New("order not found")
	}

	// 返回更新后的订单
	var order entity.ProductOrder
	err = db.DB.QueryRow("SELECT id, customer_id, amount, status, created_at FROM orders WHERE id = ?", id).Scan(
		&order.ID, &order.CustomerID, &order.Amount, &order.Status, &order.CreatedAt,
	)
	if err != nil {
		return entity.ProductOrder{}, err
	}

	return order, nil
}
