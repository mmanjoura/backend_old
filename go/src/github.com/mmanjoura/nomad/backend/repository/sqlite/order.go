package sqlite

import (
	"context"
	"strings"

	"github.com/mmanjoura/nomad/backend"
	order "github.com/mmanjoura/nomad/backend/order"
)

var _ order.OrderService = (*OrderService)(nil)

type OrderService struct {
	db *DB
}

func NewOrderService(db *DB) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) FindOne(ctx context.Context, id int) (*order.Order, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Fetch order and their associated OAuth objects.
	order, err := findOrderByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) FindAll(ctx context.Context, orderId int, filter backend.Filter) ([]*order.Order, int, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback()
	return findOrders(ctx, tx, orderId, filter)
}

func (s *OrderService) Create(ctx context.Context, orderId int, order *order.Order) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create a new order object and attach associated OAuth objects.
	if err := createOrder(ctx, tx, orderId, order); err != nil {
		return err
	}
	return tx.Commit()
}

func (s *OrderService) Update(ctx context.Context, id int, c order.Order) (*order.Order, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Update order & attach associated OAuth objects.
	order, err := updateOrder(ctx, tx, id, c)
	if err != nil {
		return order, err
	} else if err := tx.Commit(); err != nil {
		return order, err
	}
	return order, nil
}

func (s *OrderService) Delete(ctx context.Context, id int) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := deleteOrder(ctx, tx, id); err != nil {
		return err
	}
	return tx.Commit()
}

func findOrderByID(ctx context.Context, tx *Tx, orderId int) (*order.Order, error) {
	a, _, err := findOrders(ctx, tx, orderId, backend.Filter{ID: &orderId})
	if err != nil {
		return nil, err
	} else if len(a) == 0 {
		return nil, &backend.Error{Code: backend.ENOTFOUND, Message: "Order not found."}
	}
	return a[0], nil
}

func findOrders(ctx context.Context, tx *Tx, orderId int, filter backend.Filter) (_ []*order.Order, n int, err error) {

	// Build WHERE clause.
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}

	// Execute query to fetch order childeren values rows.
	rows, err := tx.QueryContext(ctx, GetOrders+strings.Join(where, " AND ")+`
		ORDER BY customer_order.id ASC
		`+FormatLimitOffset(filter.Limit, filter.Offset),
		args...,
	)
	if err != nil {
		return nil, n, err
	}
	defer rows.Close()

	orders := make([]*order.Order, 0)
	for rows.Next() {
		var custorder order.Order
		var shippingAddress order.ShippingAddress
		var billingAddress order.BillingAddress
		var customer order.Customer
		var orderStatus order.Status

		if err := rows.Scan(
			&custorder.ID,
			&custorder.TrackingNumber,
			&custorder.CustomerID,
			&custorder.CustomerContact,
			&custorder.Amount,
			&custorder.SalesTax,
			&custorder.PaidTotal,
			&custorder.Total,
			&custorder.CouponID,
			&custorder.ParentID,
			&custorder.ShopID,
			&custorder.Discount,
			&custorder.PaymentID,
			&custorder.PaymentGateway,
			&custorder.LogisticsProvider,
			&custorder.DeliveryFee,
			(*NullTime)(&custorder.DeliveryTime),
			(*NullTime)(&custorder.DeletedAt),
			(*NullTime)(&custorder.CreatedAt),
			(*NullTime)(&custorder.UpdatedAt),
			&shippingAddress.Zip,
			&shippingAddress.City,
			&shippingAddress.State,
			&shippingAddress.Country,
			&shippingAddress.StreetAddress,
			&billingAddress.Zip,
			&billingAddress.City,
			&billingAddress.State,
			&billingAddress.Country,
			&billingAddress.StreetAddress,
			&customer.ID,
			&customer.Name,
			&customer.Email,
			(*NullTime)(&customer.EmailVerifiedAt),
			(*NullTime)(&customer.CreatedAt),
			(*NullTime)(&customer.UpdatedAt),
			&customer.IsActive,
			&customer.ShopID,
			&orderStatus.ID,
			&orderStatus.Name,
			&orderStatus.Serial,
			&orderStatus.Color,
			(*NullTime)(&orderStatus.CreatedAt),
			(*NullTime)(&orderStatus.UpdatedAt),
		); err != nil {
			return nil, 0, err
		}

		orders = append(orders, &custorder)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return orders, n, nil
}

func createOrder(ctx context.Context, tx *Tx, orderId int, order *order.Order) error {
	// Set timestamps to the current time.
	order.CreatedAt = tx.now
	order.UpdatedAt = order.CreatedAt

	//Get this from Db
	order.ID = orderId

	// Perform basic field validation.
	if err := order.Validate(); err != nil {
		return err
	}

	// Execute insertion query.
	result, err := tx.ExecContext(ctx, `
		INSERT INTO order (
			Shop_id,
			slug,
			name,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`,
		order.Amount,
		order.CouponID,
		order.ID,
		(*NullTime)(&order.CreatedAt),
		(*NullTime)(&order.UpdatedAt),
	)
	if err != nil {
		return FormatError(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	order.ID = int(id)

	return nil
}

func updateOrder(ctx context.Context, tx *Tx, id int, attr order.Order) (*order.Order, error) {
	// Fetch current object state.
	order, err := findOrderByID(ctx, tx, id)
	if err != nil {
		return order, err
	} //else if order.ID != order.AttributeIDFromContext(ctx) {
	// 	return nil, nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to update this order.")
	// }

	// Update fields.
	if v := attr.PaymentGateway; v != "" {
		order.PaymentGateway = v
	}
	if v := attr.PaymentGateway; v != "" {
		order.PaymentGateway = v
	}

	// Set last updated date to current time.
	order.UpdatedAt = tx.now

	// Perform basic field validation.
	if err := order.Validate(); err != nil {
		return order, err
	}

	// Execute update query.
	if _, err := tx.ExecContext(ctx, `
		UPDATE order
		SET slug = ?,
		    name = ?,
		    updated_at = ?
		WHERE id = ?
	`,
		order.PaymentGateway,
		order.PaymentGateway,
		(*NullTime)(&order.UpdatedAt),
		id,
	); err != nil {
		return order, FormatError(err)
	}

	return order, nil
}

func deleteOrder(ctx context.Context, tx *Tx, orderId int) error {
	// Verify object exists.
	if _, err := findUserByID(ctx, tx, orderId); err != nil {
		return err
	} //else if user.ID != user.UserIDFromContext(ctx) {
	// 	return nomad.Errorf(nomad.EUNAUTHORIZED, "You are not allowed to delete this order.")
	// }

	// Remove row from database.
	if _, err := tx.ExecContext(ctx, `DELETE FROM order WHERE id = ?`, orderId); err != nil {
		return FormatError(err)
	}
	return nil
}
