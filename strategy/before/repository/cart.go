package repository

type (
	cart struct {
	}

	ICartRepo interface {
	}
)

func NewCartRepo() ICartRepo {
	return &cart{}
}
