package service

import (
	"fmt"
	"time"

	"github.com/cikupin/testing-github-action/repository"
)

type Svc struct {
	repo *repository.Repository
}

func NewSvc(r *repository.Repository) *Svc {
	return &Svc{
		repo: r,
	}
}

func (s *Svc) Process() error {
	fmt.Println("[Svc.Process] Inside process service, additional process takes time 785ms")
	time.Sleep(785 * time.Millisecond)

	s.repo.Get()
	return nil
}

func (s *Svc) Insert() error {
	fmt.Println("[Svc.Insert] Inside process service, additional process takes time 500ms")
	time.Sleep(500 * time.Millisecond)

	s.repo.Set()
	return nil
}
