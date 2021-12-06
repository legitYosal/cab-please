package surgegrpc

import (
	context "context"
	"log"

	utils "github.com/usefss/cab-please/identity/utils"

	"google.golang.org/grpc"
)

var surgePort string = utils.GetEnv("SURGE_SERVICE_PORT")
var surgeHost string = utils.GetEnv("SURGE_SERVICE_HOST")

func GetSurgeRating(demandAreaKey string) float32 {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(surgeHost+":"+surgePort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := NewSurgeServiceClient(conn)
	res, err := c.GetSurgeRate(context.Background(), &DemandAreaKey{Key: demandAreaKey})
	if err != nil {
		log.Fatalf("Error when calling GetSurgeRate: %s", err)
	}
	log.Printf("Response from server: %f", res.Rate)
	return res.Rate
}
