package update_status

import (
	"context"
	"fmt"
	"github.com/verryp/tech-talk-lp/strategy/after/model"
	"github.com/verryp/tech-talk-lp/strategy/after/payload"
	"github.com/verryp/tech-talk-lp/strategy/after/repository"
)

type completed struct {
	shippingRepo repository.IShipping
}

func NewCompletedUpdateStatus(
	shippingRepo repository.IShipping,
) IOrderStatus {
	return &completed{
		shippingRepo: shippingRepo,
	}
}

func (c *completed) Type() string {
	return "completed"
}

func (c *completed) Update(ctx context.Context, req payload.OrderReq, order model.Order) (string, error) {
	fmt.Println("execute update status", c.Type())

	// todo: logic validate order status

	// todo: logic check shipping status

	// todo: logic update shipping status

	return c.Type(), nil
}
