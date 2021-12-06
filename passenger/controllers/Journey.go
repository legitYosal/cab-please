package controllers

import (
	"net/http"

	"github.com/usefss/cab-please/identity/permission"
	"github.com/usefss/cab-please/passenger/mapgrpc"
	"github.com/usefss/cab-please/passenger/rabbitmq"
	"github.com/usefss/cab-please/passenger/redis"
	"github.com/usefss/cab-please/passenger/surgegrpc"

	"github.com/gin-gonic/gin"
)

type JourneyRequest struct {
	StartLatitude  float64 `json:"start_latitude" binding:"required"`
	StartLongitude float64 `json:"start_longitude" binding:"required"`
}

var prefixKey string = "surge-demands:"

func adjustJourneyDemands(districtKey string) {
	// Note I know I should prevent one user to send many demands but it actually \
	// must be handled in ratelimit
	redis.IncrementOrSetDemand(prefixKey + districtKey)
	rabbitmq.PublishDecreaseDemand(prefixKey + districtKey)
}

func getSurgeRating(districtKey string) float32 {
	return surgegrpc.GetSurgeRating(districtKey)
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
	_, isAuthorized := permission.IsAuthenticated(c)
	if isAuthorized == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access!"})
		return
	}

	var input JourneyRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	district, city, country := mapgrpc.ResolveCoordinates(input.StartLatitude, input.StartLongitude)
	districtKey := country + ":" + city + ":" + district
	go adjustJourneyDemands(districtKey)
	surgeRating := getSurgeRating(districtKey)
	c.JSON(http.StatusOK, gin.H{
		"surge_rating": surgeRating,
	})
}
