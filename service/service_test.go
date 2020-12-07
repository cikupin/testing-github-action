package service_test

import (
	"testing"

	"github.com/cikupin/testing-github-action/repository"
	"github.com/cikupin/testing-github-action/service"
	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	repo := repository.NewRepository()
	svc := service.NewSvc(repo)

	result := svc.Process()
	assert.Nil(t, result)
}

func TestInsert(t *testing.T) {
	repo := repository.NewRepository()
	svc := service.NewSvc(repo)

	result := svc.Insert()
	assert.Nil(t, result)
}
