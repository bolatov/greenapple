package main

// http://thenewstack.io/make-a-restful-json-api-go/
import (
	"flag"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	portPtr := flag.String("port", "8080", "port number")
	flag.Parse()

	log.Fatal(http.ListenAndServe(":"+(*portPtr), router))
}
