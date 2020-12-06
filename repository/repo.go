package repository

import (
	"fmt"
	"time"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Get() {
	fmt.Println("[Repo.Get] add 2 seconds delay...")
	time.Sleep(2 * time.Second)
	return
}
