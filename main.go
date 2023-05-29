package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: check_my_health [host] [port]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Missing some params!")
		os.Exit(1)
	}

	check_tcp_port_result, err := Check_tcp_port(args[0], args[1])

	if err != nil {
		log.Fatalf("Error during TCP port checking: %s", err)
	} else {
		log.Print(check_tcp_port_result)
	}

	os.Exit(0)
}
