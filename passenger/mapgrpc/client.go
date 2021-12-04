package mapgrpc

import (
	context "context"
	"fmt"
	"log"

	utils "github.com/usefss/cab-please/identity/utils"

	"google.golang.org/grpc"
)

var mapPort string = utils.GetEnv("MAP_SERVICE_PORT")
var mapHost string = utils.GetEnv("MAP_SERVICE_HOST")

func ResolveCoordinates(latitude float64, longitude float64) (string, string, string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+mapPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := NewMapClient(conn)
	response, err := c.GetServerResponse(
		context.Background(),
		&CoordinateRequest{
			Latitude:  fmt.Sprintf("%v", latitude),
			Longitude: fmt.Sprintf("%v", longitude),
		})
	if err != nil {
		log.Fatalf("Error when calling get server response: %s", err)
	}
	return response.District, response.City, response.Country
}
