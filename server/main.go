package main

import (
	"fmt"
	"os"
)

const mongoPath = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

func main() {
	signalHandler()

	switch len(os.Args) {
	case 1:
		startTcp()
	case 2:
		if os.Args[1] == "http" {
			startHttp()
			break
		}
		fallthrough
	default:
		fmt.Println("use \"search_server\" without arguments to start simple tcp-server for search by phone or \"search_server http\" for http-server")
	}
}
