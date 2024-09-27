package main

import (
	"fmt"
	"github.com/kobradag/kobrad/cmd/kobrawallet/libkobrawallet"
	"github.com/kobradag/kobrad/util"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		panic(err)
	}

	privateKey, publicKey, err := libkobrawallet.CreateKeyPair(false)
	if err != nil {
		panic(err)
	}

	addr, err := util.NewAddressPublicKey(publicKey, cfg.NetParams().Prefix)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Private key: %x\n", privateKey)
	fmt.Printf("Address: %s\n", addr)
}
