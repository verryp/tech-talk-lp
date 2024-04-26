package usecase

import (
	"context"
	"fmt"

	"github.com/verryp/tech-talk-lp/strategy/after/payload"
	"github.com/verryp/tech-talk-lp/strategy/after/repository"
	"github.com/verryp/tech-talk-lp/strategy/after/usecase/update_status"
)

type (
	order struct {
		orderRepo           repository.IOrderRepo
		updateOrderStrategy update_status.Strategy
	}

	IOrderUseCase interface {
		UpdateStatus(id string, req payload.OrderReq) error
	}
)

func NewOrderUseCase(
	orderRepo repository.IOrderRepo,
	updateOrderStrategy ...update_status.IOrderStatus,
) IOrderUseCase {
	return &order{
		orderRepo:           orderRepo,
		updateOrderStrategy: update_status.Registry(updateOrderStrategy...),
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

	resolver := o.updateOrderStrategy.Get(req.Status)
	if resolver == nil {
		err = fmt.Errorf("invalid status")
		return
	}

	orderStatus, err := resolver.Update(context.Background(), req, *ord)
	if err != nil {
		return
	}

	ord.Status = orderStatus

	return o.orderRepo.Update(*ord)

}
