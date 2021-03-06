package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cikupin/testing-github-action/service"
)

type Response struct {
	Status       string `json:"status"`
	Message      string `json:"message,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type Hdlr struct {
	svc *service.Svc
}

func NewHdlr(s *service.Svc) *Hdlr {
	return &Hdlr{
		svc: s,
	}
}

func (h *Hdlr) GetData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handler.GetData] Going to handler, additional process 400ms")
	time.Sleep(400 * time.Millisecond)

	h.svc.Process()
	fmt.Println()

	resp := Response{
		Status:  "OK",
		Message: "data successfully being retrieved",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Hdlr) SetData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handler.SetData] Going to handler, additional process 400ms")
	time.Sleep(400 * time.Millisecond)

	h.svc.Process()
	fmt.Println()

	resp := Response{
		Status:  "OK",
		Message: "data successfully being inserted",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Hdlr) ErrorData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handler.ErrorData] Going to handler, additional process 400ms")
	log.Println("[Handler.ErrorData] ERROR happened here!")
	log.Println("[Handler.ErrorData] Now exiting process!")
	fmt.Println()

	resp := Response{
		Status:       "ERROR",
		ErrorMessage: "Opps, something wrong happened",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(resp)
}
