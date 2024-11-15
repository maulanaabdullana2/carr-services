package order

import (
	model "cart-order-service/repository/models"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{db}
}

func (o *store) CreateOrder(bReq model.Order) (*uuid.UUID, *string, error) {
	tx, err := o.db.Begin()
	if err != nil {
		return nil, nil, err
	}

	query := `
	INSERT INTO orders (
		user_id,
			payment_type_id,
			order_number,
			total_price,
			product_order,
			status,
			is_paid,
			ref_code,
			created_at
	)VALUES (
		$1,$2,$3,$4,$5,$6,$7,$8,NOW()
	)RETURNING id, ref_code
	`
	var orderID uuid.UUID
	var refCode string

	if err := tx.QueryRow(
		query,
		bReq.UserID,
		bReq.PaymentTypeID,
		bReq.OrderNumber,
		bReq.TotalPrice,
		bReq.ProductOrder,
		bReq.Status,
		bReq.IsPaid,
		bReq.RefCode,
	).Scan(&orderID, &refCode); err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	return &orderID, &refCode, nil

}

func (o *store) CreateOrderItemsLogs(bReq model.OrderItemsLogs) (*string, error) {
	tx, err := o.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `
	INSERT INTO order_status_logs (
			order_id,
			ref_code,
			from_status,
			to_status,
			notes,
			created_at
		) VALUES (
			$1, $2, $3, $4, $5, NOW()
		) RETURNING ref_code
	`

	var refCode string

	if err := tx.QueryRow(
		query,
		bReq.OrderID,
		bReq.RefCode,
		bReq.FromStatus,
		bReq.ToStatus,
		bReq.Notes,
	).Scan(&refCode); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return &refCode, nil
}

func (o *store) GetOrderStatus(userID uuid.UUID, orderID uuid.UUID) (*model.Order, error) {
	query := `
		SELECT 
			user_id,
			payment_type_id,
			order_number,
			total_price,
			product_order,
			status,
			is_paid,
			ref_code,
			created_at
		FROM orders
		WHERE user_id=$1 AND id = $2
	`

	var order model.Order

	if err := o.db.QueryRow(query, userID, orderID).Scan(
		&order.UserID,
		&order.PaymentTypeID,
		&order.OrderNumber,
		&order.TotalPrice,
		&order.ProductOrder,
		&order.Status,
		&order.IsPaid,
		&order.RefCode,
		&order.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *store) UpdateStatus(req model.UpdateStatus) error {
	tx, err := o.db.Begin()
	if err != nil {
		return err
	}

	query := `
		UPDATE orders SET 
			status = $1,
			updated_at = NOW()
		WHERE user_id = $2 AND id = $3
	`
	if _, err := tx.Exec(query, req.Status, req.UserID, req.OrderID); err != nil {
		tx.Rollback()
		return errors.New("failed to update order status")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errors.New("failed to commit transaction")
	}

	return nil

}
