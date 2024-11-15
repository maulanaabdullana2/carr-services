package order

import (
	model "cart-order-service/repository/models"

	"github.com/google/uuid"
)

type orderStore interface {
	CreateOrder(bReq model.Order) (*uuid.UUID, *string, error)
	CreateOrderItemsLogs(bReq model.OrderItemsLogs) (*string, error)
	GetOrderStatus(userID uuid.UUID, orderID uuid.UUID) (*model.Order, error)
	UpdateStatus(req model.UpdateStatus) error
}

type order struct {
	store orderStore
}

func NewOrder(store orderStore) *order {
	return &order{store}
}

func (o *order) CreateOrder(bReq model.Order) (*uuid.UUID, error) {
	orderID, refCode, err := o.store.CreateOrder(bReq)
	if err != nil {
		return nil, err
	}

	_, err = o.store.CreateOrderItemsLogs(model.OrderItemsLogs{
		OrderID:    *orderID,
		RefCode:    *refCode,
		FromStatus: "",
		ToStatus:   model.OrderStatusPending,
		Notes:      "Order created",
	})
	if err != nil {
		return nil, err
	}

	return orderID, nil
}

func (o *order) GetOrderStatus(userID uuid.UUID, orderID uuid.UUID) (*model.Order, error) {
	return o.store.GetOrderStatus(userID, orderID)
}

func (o *order) UpdateStatus(req model.UpdateStatus) (*string, error) {
	successMessage := "Update status success"
	if err := o.store.UpdateStatus(req); err != nil {
		return nil, err
	}

	return &successMessage, nil
}
