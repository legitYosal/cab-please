package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/usefss/cab-please/identity/permission"
	"github.com/usefss/cab-please/passenger/protob"

	"github.com/gin-gonic/gin"
)

type JourneyRequest struct {
	StartLatitude  float64 `json:"start_latitude" binding:"required"`
	StartLongitude float64 `json:"start_longitude" binding:"required"`
}

func adjustJourneyDemands() {
	time.Sleep(time.Second * 1)
	fmt.Println("test")
}

func getClientDistrict() string {
	return "test"
}

func getSurgeRating() float64 {
	return 0.5
}

// RequestJourney godoc
// @Summary request a ride
// @Accept json
// @Produce json
// @tags Journey
// @Param journey body JourneyRequest true "Request Journey"
// @Success 200 {object} object
// @Security ApiKeyAuth
// @Router /api/journey/ [post]
func RequestJourney(c *gin.Context) {
	user, isAuthorized := permission.IsAuthenticated(c)
	if isAuthorized == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access!"})
		return
	}

	var input JourneyRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	district, city, country := protob.ResolveCoordinates(35.692437, 51.320802)
	districtKey := country + ":" + city + ":" + district
	fmt.Println(districtKey)
	go adjustJourneyDemands()
	surgeRating := getSurgeRating()
	fmt.Println((*user).ID)
	c.JSON(http.StatusOK, gin.H{
		"surge_rating": surgeRating,
	})
}
