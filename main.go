package main

import (
	"fmt"
	. "github.com/tendermint/go-common"
	"github.com/tendermint/tmsp/server"
	"github.com/tigwyk/verdis/app"
)

func main() {
	fmt.Println("test")
	_, err := server.NewServer("tcp://0.0.0.0:46658", "grpc", app.NewDummyApplication())
	if err != nil {
		panic(err)
	}
	// Wait forever
	TrapSignal(func() {
		// Cleanup
	}
}
