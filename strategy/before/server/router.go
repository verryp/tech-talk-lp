package server

import (
	"github.com/gorilla/mux"
	"github.com/verryp/tech-talk-lp/strategy/before/handler"
	"github.com/verryp/tech-talk-lp/strategy/before/repository"
	"github.com/verryp/tech-talk-lp/strategy/before/usecase"
	"log"
	"net/http"
)

func Router() {
	router := mux.NewRouter()

	orderUseCase := usecase.NewOrderUseCase(
		repository.NewOrderRepo(),
		repository.NewItemRepo(),
		repository.NewPaymentRepo(),
		repository.NewShippingRepo(),
		repository.NewPaymentRepo(),
	)

	updateOrder := handler.NewOrderHandler(orderUseCase)

	router.
		HandleFunc("/orders/{id}", updateOrder.Update).
		Methods("PUT")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}
