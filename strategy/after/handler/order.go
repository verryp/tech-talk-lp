package handler

import (
	"encoding/json"
	"net/http"

	"github.com/verryp/tech-talk-lp/strategy/after/payload"
	"github.com/verryp/tech-talk-lp/strategy/after/usecase"

	"github.com/gorilla/mux"
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

	var req payload.OrderReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.orderUseCase.UpdateStatus(id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("ok"))
}
