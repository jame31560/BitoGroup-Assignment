package main

import (
	"tinder/api"
	"tinder/docs"
	"tinder/repository"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	repo := repository.NewRepository()
	api := api.NewApi(repo)

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/add_user", api.AddSinglePersonAndMatch)
		apiGroup.DELETE("/remove_user", api.RemoveSinglePerson)
		apiGroup.GET("/query_single_user", api.QuerySinglePeople)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
