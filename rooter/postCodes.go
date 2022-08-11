package rooter

import (
	"PostCodeProject/controller"
	"github.com/gin-gonic/gin"
)

func PostCodesRoot(api *gin.RouterGroup) {
	api.GET("/cities", controller.GETCities)
	api.GET("/city/:city/counties", controller.GETCitiesCounty)
	api.GET("/city/:city/county/:county/towns", controller.GETCountiesTown)
	api.GET("/county/:county/town/:town/districts", controller.GETTownsDistinct)
}
