package repository

import (
	"fmt"
	"time"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Get() error {
	fmt.Println("[Repo.Get] add 2 seconds delay...")
	time.Sleep(2 * time.Second)
	return nil
}

func (r *Repository) Set() error {
	fmt.Println("[Repo.Set] add 1 seconds delay...")
	time.Sleep(1 * time.Second)
	return nil
}
