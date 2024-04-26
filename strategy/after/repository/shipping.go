package repository

type (
	shipping struct {
	}

	IShipping interface {
	}
)

func NewShippingRepo() IShipping {
	return &shipping{}
}
