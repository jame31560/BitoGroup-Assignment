package api

import "github.com/gin-gonic/gin"

type ApiInterface interface {
	AddSinglePersonAndMatch(c *gin.Context)
	RemoveSinglePerson(c *gin.Context)
	QuerySinglePeople(c *gin.Context)
}
