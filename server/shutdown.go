package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handleShutdown(s *Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	fmt.Println("Shutting down...")
	s.server.Close()
	fmt.Println("Shutdown complete")
}
