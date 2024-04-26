package update_status

import (
	"context"
	"github.com/verryp/tech-talk-lp/strategy/after/model"
	"github.com/verryp/tech-talk-lp/strategy/after/payload"
)

type (
	IOrderStatus interface {
		Type() string
		Update(ctx context.Context, req payload.OrderReq, order model.Order) (orderStatus string, err error)
	}
)
