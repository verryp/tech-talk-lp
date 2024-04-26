package repository

import "github.com/verryp/tech-talk-lp/strategy/before/model"

type (
	orderRepo struct {
	}

	IOrderRepo interface {
		Update(order model.Order) error
		Get(id string) (*model.Order, error)
		Create(order model.Order) error
	}
)

func (o *orderRepo) Update(order model.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepo) Get(id string) (*model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepo) Create(order model.Order) error {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepo() IOrderRepo {
	return &orderRepo{}
}
