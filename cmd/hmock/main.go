package main

import (
	"github.com/serber/hmock/internal/config"
	"github.com/serber/hmock/internal/server"

	"github.com/serber/hmock/internal/server/runner"
)

func main() {
	var portNumber uint16 = 443
	s := server.New(portNumber)

	runner.Run(s)

	config.Print()
}
