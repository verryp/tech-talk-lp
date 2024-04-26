package repository

type (
	payment struct {
	}

	IPaymentRepo interface {
	}
)

func NewPaymentRepo() IPaymentRepo {
	return &payment{}
}
