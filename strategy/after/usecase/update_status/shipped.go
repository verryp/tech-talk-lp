package update_status

import (
	"context"
	"fmt"
	"github.com/verryp/tech-talk-lp/strategy/after/model"
	"github.com/verryp/tech-talk-lp/strategy/after/payload"
	"github.com/verryp/tech-talk-lp/strategy/after/repository"
)

type shipped struct {
	paymentRepo  repository.IPaymentRepo
	shippingRepo repository.IShipping
}

func NewShippedOrderStatus(
	paymentRepo repository.IPaymentRepo,
	shippingRepo repository.IShipping,
) IOrderStatus {
	return &shipped{
		paymentRepo:  paymentRepo,
		shippingRepo: shippingRepo,
	}
}

func (s *shipped) Type() string {
	return "shipped"
}

func (s *shipped) Update(ctx context.Context, req payload.OrderReq, order model.Order) (orderStatus string, err error) {
	fmt.Println("execute update status", s.Type())

	// todo: logic validate status order

	// todo: logic check status payment

	// todo: logic create shipping

	// todo: logic refund mechanism if shipping failed to create

	return s.Type(), nil
}
