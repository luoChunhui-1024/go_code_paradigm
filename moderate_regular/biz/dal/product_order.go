package dal

import (
	"fmt"
	"github.com/go_code_paradigm/moderate_regular/biz/entity"
	"github.com/go_code_paradigm/moderate_regular/pkg/db"
	"strings"
)

// GetOrders 多条件查询订单
func GetOrders(customerIDs []int, orderIDs []int, minAmount, maxAmount float64, startDate, endDate string) ([]entity.ProductOrder, error) {
	var conditions []string
	var params []interface{}

	if len(customerIDs) > 0 {
		conditions = append(conditions, fmt.Sprintf("customer_id IN (%s)", placeholders(len(customerIDs))))
		for _, id := range customerIDs {
			params = append(params, id)
		}
	}

	if len(orderIDs) > 0 {
		conditions = append(conditions, fmt.Sprintf("id IN (%s)", placeholders(len(orderIDs))))
		for _, id := range orderIDs {
			params = append(params, id)
		}
	}

	if minAmount > 0 {
		conditions = append(conditions, "amount >= ?")
		params = append(params, minAmount)
	}

	if maxAmount > 0 {
		conditions = append(conditions, "amount <= ?")
		params = append(params, maxAmount)
	}

	if startDate != "" {
		conditions = append(conditions, "created_at >= ?")
		params = append(params, startDate)
	}

	if endDate != "" {
		conditions = append(conditions, "created_at <= ?")
		params = append(params, endDate)
	}

	query := "SELECT id, customer_id, product_id, amount, status, created_at, updated_at, extensions FROM orders"
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := db.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.ProductOrder
	for rows.Next() {
		var order entity.ProductOrder
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.ProductID, &order.Amount, &order.Status, &order.CreatedAt, &order.UpdatedAt, &order.Extensions); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// UpdateOrder 更新订单状态和扩展字段
func UpdateOrder(id int, status, extensions string) (entity.ProductOrder, error) {
	result, err := db.DB.Exec("UPDATE orders SET status = ?, extensions = ?, updated_at = NOW() WHERE id = ?", status, extensions, id)
	if err != nil {
		return entity.ProductOrder{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return entity.ProductOrder{}, err
	}

	if rowsAffected == 0 {
		return entity.ProductOrder{}, fmt.Errorf("order not found")
	}

	// 返回更新后的订单
	var order entity.ProductOrder
	err = db.DB.QueryRow("SELECT id, customer_id, product_id, amount, status, created_at, updated_at, extensions FROM orders WHERE id = ?", id).Scan(
		&order.ID, &order.CustomerID, &order.ProductID, &order.Amount, &order.Status, &order.CreatedAt, &order.UpdatedAt, &order.Extensions,
	)
	if err != nil {
		return entity.ProductOrder{}, err
	}

	return order, nil
}

// 辅助函数：生成 SQL IN 子句的占位符
func placeholders(count int) string {
	return strings.TrimRight(strings.Repeat("?,", count), ",")
}
