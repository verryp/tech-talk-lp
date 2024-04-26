package update_status

import (
	"context"
	"fmt"
	"github.com/verryp/tech-talk-lp/strategy/after/model"
	"github.com/verryp/tech-talk-lp/strategy/after/payload"
	"github.com/verryp/tech-talk-lp/strategy/after/repository"
)

type pending struct {
	itemRepo    repository.IItemRepo
	cartRepo    repository.ICartRepo
	paymentRepo repository.IPaymentRepo
}

func NewPendingUpdateStatus(
	itemRepo repository.IItemRepo,
	cartRepo repository.ICartRepo,
	paymentRepo repository.IPaymentRepo,
) IOrderStatus {
	return &pending{
		itemRepo:    itemRepo,
		cartRepo:    cartRepo,
		paymentRepo: paymentRepo,
	}
}

func (p *pending) Type() string {
	return "pending"
}

func (p *pending) Update(ctx context.Context, req payload.OrderReq, order model.Order) (orderStatus string, err error) {
	fmt.Println("execute update status", p.Type())

	// todo: logic to validate the status order must be "created", if else return error

	// todo: logic get item

	// todo: logic check stock

	// todo: logic validate total amount

	// todo: logic remove cart

	// todo: logic create payment with status pending

	// todo: logic deduct item stock

	// todo: rollback mechanism

	return p.Type(), nil
}
