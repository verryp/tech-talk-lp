package update_status

import (
	"context"
	"fmt"
	"github.com/verryp/tech-talk-lp/strategy/after/model"
	"github.com/verryp/tech-talk-lp/strategy/after/payload"
	"github.com/verryp/tech-talk-lp/strategy/after/repository"
)

type paid struct {
	paymentRepo repository.IPaymentRepo
}

func NewPaidUpdateStatus(
	paymentRepo repository.IPaymentRepo,
) IOrderStatus {
	return &paid{
		paymentRepo: paymentRepo,
	}
}

func (p *paid) Type() string {
	return "paid"
}

func (p *paid) Update(ctx context.Context, req payload.OrderReq, order model.Order) (orderStatus string, err error) {
	fmt.Println("execute update status", p.Type())

	// todo: logic to validate the status order must be "pending", if else return error

	// todo: logic check status payment if already paid or not

	// todo: logic if payment not paid, then return the stock item and cancel cancel the order

	// todo: logic update payment status to paid

	// todo: logic rollback mechanism

	return p.Type(), nil
}
