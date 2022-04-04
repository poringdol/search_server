package http

import (
	"fmt"
	"net/http"
)

const serverAddr = "localhost:9000"
const mongoPath = "mongodb://localhost:27017/?maxPoolSize=20&w=majority"

func StartHttp() {
	http.HandleFunc("/find_by_phone", findByPhone)
	http.HandleFunc("/find_by_address", findByAddress)
	http.HandleFunc("/find_by_name", findByName)
	http.HandleFunc("/get_cities", getCities)
	http.HandleFunc("/", root)

	fmt.Printf("HTTP-server started at %s\n", serverAddr)
	http.ListenAndServe(serverAddr, nil)
}
