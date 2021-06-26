package main

import (
	"net"
	"os"

	protos "github.com/cesarvspr/grpc-currency/protos/currency"
	"github.com/cesarvspr/grpc-currency/server"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()

	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("unable to listen", "error", err)
		os.Exit(1)
	} else {
		log.Info("RUNNING")
	}

	gs.Serve(l)
}
