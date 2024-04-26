package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/verryp/tech-talk-lp/strategy/before/model"
	"github.com/verryp/tech-talk-lp/strategy/before/payload"
	"github.com/verryp/tech-talk-lp/strategy/before/usecase"
	"net/http"
)

type Order struct {
	orderUseCase usecase.IOrderUseCase
}

func NewOrderHandler(ord usecase.IOrderUseCase) *Order {
	return &Order{
		orderUseCase: ord,
	}
}

func (h *Order) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	var req model.Order
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.orderUseCase.UpdateStatus(id, payload.OrderReq{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("ok"))
}
