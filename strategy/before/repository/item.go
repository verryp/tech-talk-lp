package repository

type (
	item struct {
	}

	IItemRepo interface {
	}
)

func NewItemRepo() IItemRepo {
	return &item{}
}
