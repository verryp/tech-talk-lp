package repository

import "github.com/verryp/tech-talk-lp/strategy/after/model"

type (
	orderRepo struct {
	}

	IOrderRepo interface {
		Update(order model.Order) error
		Get(id string) (*model.Order, error)
		Create(order model.Order) error
	}
)

func NewOrderRepo() IOrderRepo {
	return &orderRepo{}
}

func (o *orderRepo) Update(order model.Order) error {
	return nil
}

func (o *orderRepo) Get(id string) (*model.Order, error) {
	return &model.Order{}, nil
}

func (o *orderRepo) Create(order model.Order) error {
	return nil
}
