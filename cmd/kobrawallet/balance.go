package main

import (
	"context"
	"fmt"

	"github.com/kobradag/kobrad/cmd/kobrawallet/daemon/client"
	"github.com/kobradag/kobrad/cmd/kobrawallet/daemon/pb"
	"github.com/kobradag/kobrad/cmd/kobrawallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatKobra(addressBalance.Available), utils.FormatKobra(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, KODA %s %s%s\n", utils.FormatKobra(response.Available), utils.FormatKobra(response.Pending), pendingSuffix)

	return nil
}
