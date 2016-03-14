package main

// http://thenewstack.io/make-a-restful-json-api-go/
import (
	"flag"
	"log"
	"net/http"
	"os"
)

func init() {
	f, err := os.OpenFile("algo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
}

func main() {
	portPtr := flag.String("port", "8080", "port number")
	flag.Parse()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+(*portPtr), router))
}
