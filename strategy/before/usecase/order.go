package usecase

import (
	"fmt"
	"github.com/verryp/tech-talk-lp/strategy/before/payload"
	"github.com/verryp/tech-talk-lp/strategy/before/repository"
)

type (
	order struct {
		orderRepo    repository.IOrderRepo
		itemRepo     repository.IItemRepo
		cartRepo     repository.ICartRepo
		paymentRepo  repository.IPaymentRepo
		shippingRepo repository.IShipping
	}

	IOrderUseCase interface {
		UpdateStatus(id string, req payload.OrderReq) error
	}
)

func NewOrderUseCase(
	orderRepo repository.IOrderRepo,
	itemRepo repository.IItemRepo,
	cartRepo repository.ICartRepo,
	paymentRepo repository.IPaymentRepo,
	shippingRepo repository.IShipping,
) IOrderUseCase {
	return &order{
		orderRepo:    orderRepo,
		itemRepo:     itemRepo,
		cartRepo:     cartRepo,
		paymentRepo:  paymentRepo,
		shippingRepo: shippingRepo,
	}
}

func (o *order) UpdateStatus(id string, req payload.OrderReq) (err error) {
	ord, err := o.orderRepo.Get(id)
	if err != nil {
		return
	}

	if ord == nil {
		err = fmt.Errorf("order not found")
		return
	}

	switch req.Status {
	case "pending":
		// todo: logic to validate the status order must be "created", if else return error

		// todo: logic get item

		// todo: logic check stock

		// todo: logic validate total amount

		// todo: logic remove cart

		// todo: logic create payment with status pending

		// todo: logic deduct item stock

		// todo: rollback mechanism

		ord.Status = "pending"
	case "paid":
		// todo: logic to validate the status order must be "pending", if else return error

		// todo: logic check status payment if already paid or not

		// todo: logic if payment not paid, then return the stock item and cancel cancel the order

		// todo: logic update payment status to paid

		// todo: logic rollback mechanism

		ord.Status = "paid"
	case "shipped":
		// todo: logic validate status order

		// todo: logic check status payment

		// todo: logic create shipping

		// todo: logic refund mechanism if shipping failed to create

		ord.Status = "shipped"
	case "delivered":
		// todo: logic validate status order

		// todo: logic check shipping status

		// todo: update shipping status

		// todo: logic refund mechanism if shipping has issue from driver/customer

		ord.Status = "delivered"
	case "completed":
		// todo: logic validate order status

		// todo: logic check shipping status

		// todo: logic update shipping status
		ord.Status = "completed"
	}

	return o.orderRepo.Update(*ord)

}
