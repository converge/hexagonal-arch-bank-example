package bank_handler

import (
	"fmt"
	"hexagonal-example/internal/core/ports"
	"log"
	"net/http"
)

type HTTPHandler struct {
	bankService ports.BankServiceInterface
}

func NewHTTPHandler (bankService ports.BankServiceInterface) *HTTPHandler {
	return &HTTPHandler{
		bankService: bankService,
	}
}
func (h *HTTPHandler) SendSMS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("smsSent")
	//err := h.bankService.SendSMS("ok")
	//if err != nil {
	//	log.Println(err)
	//}
	w.Write([]byte("ok"))
}

func (h *HTTPHandler) Balance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	val, err := h.bankService.Balance(1)
	if err != nil {
		log.Println()
	}
	fmt.Println(val)
	w.Write([]byte("create"))
}
