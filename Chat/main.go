package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"./chat_app/client"
	"./chat_app/server"
)

const (
	DefaultIP   = "127.0.0.1"
	DefaultPORT = "4242"
)

func usage() {
	log.Print("Unrecognized option")
	log.Println("Usage: go run main.go --mode <mode>")
	log.Print("- mode: \"server\" or \"client\"")
	log.Println("example to run a server: go run main.go --mode server")
	log.Println("example to run a client: go run main.go --mode client")
}

func options() {
	var mode string

	flag.StringVar(&mode, "mode", "client", "--mode client or --mode server")
	flag.Parse()

	if strings.ToLower(mode) == "server" {
		server, err := server.New(DefaultIP, DefaultPORT)
		if err != nil {
			log.Fatalf("Error creating server: %v", err)
		}
		server.Run()
	} else if strings.ToLower(mode) == "client" {
		client, err := client.New(DefaultIP, DefaultPORT)
		if err != nil {
			log.Fatalf("Error creating client: %v", err)
		}
		client.Run()
	} else {
		usage()
		os.Exit(2)
	}
}

func main() {
	options()
}
