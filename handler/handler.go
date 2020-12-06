package handler

import (
	"encoding/json"
	"fmt"
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

func (h *Hdlr) SetData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Handler.SetData] Going to handler, additional process 400ms")
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
