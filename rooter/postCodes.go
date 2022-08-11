package rooter

import (
	"PostCodeProject/controller"
	"github.com/gin-gonic/gin"
)

func PostCodesRoot(api *gin.RouterGroup) {
	api.GET("/cities", controller.GETCities)
	api.GET("/cities/:city/counties", controller.GETCitiesCounty)
	api.GET("/cities/:city/counties/:county/towns", controller.GETCountiesTown)
	api.GET("/cities/:city/counties/:county/towns/:town/postcode", controller.GETTownsDistinct)
}
