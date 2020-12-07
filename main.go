package main

import (
	"fmt"
	"net/http"

	"github.com/cikupin/testing-github-action/handler"
	"github.com/cikupin/testing-github-action/repository"
	"github.com/cikupin/testing-github-action/service"
	"github.com/go-chi/chi"
)

func main() {
	// dependecy injection
	repo := repository.NewRepository()
	service := service.NewSvc(repo)
	handler := handler.NewHdlr(service)

	r := chi.NewRouter()
	r.Get("/", handler.GetData)
	r.Get("/set", handler.SetData)
	r.Get("/error", handler.ErrorData)

	fmt.Println("Starting app on port 3000...")
	http.ListenAndServe(":3000", r)
}
