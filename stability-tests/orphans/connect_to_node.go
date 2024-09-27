package main

import (
	"fmt"
	"os"

	"github.com/kobradag/kobrad/infrastructure/config"
	"github.com/kobradag/kobrad/infrastructure/network/netadapter/standalone"
)

func connectToNode() *standalone.Routes {
	cfg := activeConfig()

	kobradConfig := config.DefaultConfig()
	kobradConfig.NetworkFlags = cfg.NetworkFlags

	minimalNetAdapter, err := standalone.NewMinimalNetAdapter(kobradConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating minimalNetAdapter: %+v", err)
		os.Exit(1)
	}
	routes, err := minimalNetAdapter.Connect(cfg.NodeP2PAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error connecting to node: %+v", err)
		os.Exit(1)
	}
	return routes
}
