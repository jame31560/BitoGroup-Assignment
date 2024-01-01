package api

import (
	"tinder/repository"

	"github.com/gin-gonic/gin"
)

type api struct {
	repo repository.RepositoryInterface
}

func NewApi(repo repository.RepositoryInterface) ApiInterface {
	return &api{
		repo: repo,
	}
}

func (a *api) AddSinglePersonAndMatch(c *gin.Context) {}

func (a *api) RemoveSinglePerson(c *gin.Context) {}

func (a *api) QuerySinglePeople(c *gin.Context) {}
