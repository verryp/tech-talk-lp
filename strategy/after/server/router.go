package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/verryp/tech-talk-lp/strategy/after/handler"
	"github.com/verryp/tech-talk-lp/strategy/after/repository"
	"github.com/verryp/tech-talk-lp/strategy/after/usecase"
	"github.com/verryp/tech-talk-lp/strategy/after/usecase/update_status"

	"github.com/gorilla/mux"
)

var (
	port = 8081
)

func Router() {
	router := mux.NewRouter()

	orderUseCase := usecase.NewOrderUseCase(
		repository.NewOrderRepo(),
		update_status.NewPendingUpdateStatus(repository.NewItemRepo(), repository.NewCartRepo(), repository.NewPaymentRepo()),
		update_status.NewPaidUpdateStatus(repository.NewPaymentRepo()),
		update_status.NewShippedOrderStatus(repository.NewPaymentRepo(), repository.NewShippingRepo()),
		update_status.NewDeliveredUpdateStatus(repository.NewShippingRepo()),
		update_status.NewCompletedUpdateStatus(repository.NewShippingRepo()),
	)

	updateOrder := handler.NewOrderHandler(orderUseCase)

	router.
		HandleFunc("/orders/{id}", updateOrder.Update).
		Methods(http.MethodPut)

	// Start the HTTP server
	fmt.Println("start listening on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		log.Fatal(err)
	}
}
