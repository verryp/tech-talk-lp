package update_status

import (
	"context"
	"fmt"
	"github.com/verryp/tech-talk-lp/strategy/after/model"
	"github.com/verryp/tech-talk-lp/strategy/after/payload"
	"github.com/verryp/tech-talk-lp/strategy/after/repository"
)

type delivered struct {
	shippingRepo repository.IShipping
}

func NewDeliveredUpdateStatus(
	shippingRepo repository.IShipping,
) IOrderStatus {
	return &delivered{
		shippingRepo: shippingRepo,
	}
}

func (d *delivered) Type() string {
	return "delivered"
}

func (d *delivered) Update(ctx context.Context, req payload.OrderReq, order model.Order) (orderStatus string, err error) {
	fmt.Println("execute update status", d.Type())

	// todo: logic validate status order

	// todo: logic check shipping status

	// todo: update shipping status

	// todo: logic refund mechanism if shipping has issue from driver/customer

	return d.Type(), nil
}
