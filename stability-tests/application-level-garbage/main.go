package main

import (
	"fmt"
	"os"

	"github.com/kobradag/kobrad/infrastructure/config"
	"github.com/kobradag/kobrad/infrastructure/network/netadapter/standalone"
	"github.com/kobradag/kobrad/stability-tests/common"
	"github.com/kobradag/kobrad/util/panics"
	"github.com/kobradag/kobrad/util/profiling"
)

func main() {
	defer panics.HandlePanic(log, "applicationLevelGarbage-main", nil)
	err := parseConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing config: %+v", err)
		os.Exit(1)
	}
	defer backendLog.Close()
	common.UseLogger(backendLog, log.Level())
	cfg := activeConfig()
	if cfg.Profile != "" {
		profiling.Start(cfg.Profile, log)
	}

	kobradConfig := config.DefaultConfig()
	kobradConfig.NetworkFlags = cfg.NetworkFlags

	minimalNetAdapter, err := standalone.NewMinimalNetAdapter(kobradConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating minimalNetAdapter: %+v", err)
		backendLog.Close()
		os.Exit(1)
	}

	blocksChan, err := readBlocks()
	if err != nil {
		log.Errorf("Error reading blocks: %+v", err)
		backendLog.Close()
		os.Exit(1)
	}

	err = sendBlocks(cfg.NodeP2PAddress, minimalNetAdapter, blocksChan)
	if err != nil {
		log.Errorf("Error sending blocks: %+v", err)
		backendLog.Close()
		os.Exit(1)
	}
}
