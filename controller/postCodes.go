package controller

import (
	"PostCodeProject/config"
	"PostCodeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GETCities gives all cities
func GETCities(c *gin.Context) {
	c.JSON(http.StatusOK, config.Cities)

}

// GETCitiesCounty gives you the county list of the given city
func GETCitiesCounty(c *gin.Context) {
	cityParam := c.Param("city")
	if cityParam == "" {
		c.JSON(http.StatusBadRequest, "city param is required")
		return
	}
	var returnValue []models.County
	for _, county := range config.Counties {
		if county.City == cityParam {
			returnValue = append(returnValue, county)
		}
	}

	c.JSON(http.StatusOK, returnValue)
}

// GETCountiesTown gives you the towns of the given city and county
func GETCountiesTown(c *gin.Context) {
	countyParam := c.Param("county")
	cityParam := c.Param("city")
	if countyParam == "" {
		c.JSON(http.StatusBadRequest, "county param is required")
		return
	}
	var returnValue []models.Town
	for _, town := range config.Towns {
		if town.County == countyParam && town.City == cityParam {
			returnValue = append(returnValue, town)
		}
	}
	c.JSON(http.StatusOK, returnValue)

}

// GETTownsDistinct gets post code, district, county, town and city by given county and town
func GETTownsDistinct(c *gin.Context) {
	townParam := c.Param("town")
	countyParam := c.Param("county")
	if townParam == "" {
		c.JSON(http.StatusBadRequest, "town param is required")
		return
	}
	var returnValue []models.District
	for _, district := range config.Districts {
		if district.Town == townParam && district.County == countyParam {
			returnValue = append(returnValue, district)
		}
	}
	c.JSON(http.StatusOK, returnValue)
}
