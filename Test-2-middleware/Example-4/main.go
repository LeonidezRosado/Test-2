package main

import (
	"log"
	"net/http"
	"os"
	"io"

	"github.com/gorilla/handlers"
)

//create the middleware
func theLoggingHandler(dst io.Writer) func (http.Handler) http.Handler {
	return func (h http.Handler) http.Handler  {
		return handlers.LoggingHandler(dst, h)		
	}
}

//create our handler
func OurHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"index.html")
}

func main() {
	//create our file
	logfile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND,0664) 
	if err != nil {
		log.Fatal(err)
	}

	logginghandler := theLoggingHandler(logfile)

	//create our mutliplexer/router
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(OurHandler)
	mux.Handle("/home", logginghandler(finalHandler))

	//create our server on the specifed port
	log.Print("starting on :3000")
	err = http.ListenAndServe(":3000", mux)
	log.Fatal(err)

}