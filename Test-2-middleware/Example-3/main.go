package main 

import (
	"log"
	"net/http"
	"github.com/goji/httpauth"
)

//create a basic handler that will be used for authentication using 
//third party middleware to protect it
func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r, "index.html")
}

func main() {
	//create a variable that will use the third party middleware using a function from the package
	AuthenticationHandler := httpauth.SimpleBasicAuth("Jennifer", "carrots")

	//creat out multiplexer/router
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(handler)
	mux.Handle("/home", AuthenticationHandler(finalHandler))

	//creat port to host our server
	log.Print("starting server on :4000") 
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}