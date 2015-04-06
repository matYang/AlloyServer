package main

import (
	"fmt"
	"github.com/icub3d/graceful"
	"github.com/matYang/AlloyServer/alsParser"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	alsParser.RunParser()

	// Start the server.
	fmt.Println("Using PID:", os.Getpid())

	router := NewRouter()
	log.Fatal(graceful.ListenAndServe(":8050", router))
}
