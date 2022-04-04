package runner

import (
	"github.com/serber/hmock/internal/server"
)

func Run(server *server.Server) {
	server.PrintInfo()
}
