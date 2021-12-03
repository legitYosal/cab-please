package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/usefss/cab-please/identity/permission"

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
	/*
		do following things
		 - if user is has no recent demands, +1 the demand district count and -1 message to queue
		 - get surge rating of the area
	*/
	getClientDistrict()
	go adjustJourneyDemands()
	surgeRating := getSurgeRating()
	fmt.Println((*user).ID)
	c.JSON(http.StatusOK, gin.H{
		"surge_rating": surgeRating,
	})
}
