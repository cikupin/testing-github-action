package repository_test

import (
	"testing"

	"github.com/cikupin/testing-github-action/repository"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	repo := repository.NewRepository()

	result := repo.Set()
	assert.Nil(t, result)
}

func TestGet(t *testing.T) {
	repo := repository.NewRepository()

	result := repo.Get()
	assert.Nil(t, result)
}
