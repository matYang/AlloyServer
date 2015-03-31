package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/icub3d/graceful"
)

func main() {
	// Listen for the SIGTERM.
	// Graceful shutdown
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Print("got SIGHUP, shutting down.")
		graceful.Close()
	}()

	// Start the server.
	fmt.Println("Using PID:", os.Getpid())

	router := NewRouter()
	log.Fatal(graceful.ListenAndServe(":8050", router))
}
