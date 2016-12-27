package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/tendermint/go-common"
	"github.com/tendermint/tmsp/types"
	"google.golang.org/grpc"
)

// var maxNumberConnections = 2

type GRPCServer struct {
	common.QuitService

	proto    string
	addr     string
	listener net.Listener
	server   *grpc.Server

	app types.TMSPApplicationServer
}

func NewGRPCServer(protoAddr string, app types.TMSPApplicationServer) (common.Service, error) {
	parts := strings.SplitN(protoAddr, "://", 2)
	proto, addr := parts[0], parts[1]
	s := &GRPCServer{
		proto:    proto,
		addr:     addr,
		listener: nil,
		app:      app,
	}
	s.QuitService = *common.NewQuitService(nil, "TMSPServer", s)
	_, err := s.Start() // Just start it
	return s, err
}

func (s *GRPCServer) OnStart() error {
	s.QuitService.OnStart()
	ln, err := net.Listen(s.proto, s.addr)
	if err != nil {
		return err
	}
	s.listener = ln
	s.server = grpc.NewServer()
	types.RegisterTMSPApplicationServer(s.server, s.app)
	go s.server.Serve(s.listener)
	return nil
}

func (s *GRPCServer) OnStop() {
	s.QuitService.OnStop()
	s.server.Stop()
}

func main() {
	fmt.Println("vim-go")
}
