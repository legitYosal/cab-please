package main

import (
	"fmt"

	"github.com/usefss/cab-please/surge/stubs"
	"golang.org/x/net/context"
)

type GRPCServer struct {
}

func (s *GRPCServer) GetSurgeRate(ctx context.Context, in *stubs.DemandAreaKey) (*stubs.SurgeRate, error) {
	fmt.Println(in.Key)
	return &stubs.SurgeRate{Rate: 0.78}, nil
}
