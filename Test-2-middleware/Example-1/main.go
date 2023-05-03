package main

import (
	"log"
	"net/http"
)

//creating our first middleware
func firstmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//when middleware is executed, print out a message
		log.Print("Exectuing the first middleware")
		next.ServeHTTP(w,r)//used to pass the request to the next handler

		//will print out a second message when it comes back again
		log.Print("Executing the first middleware again")
	})
}

//creating our second middleware
func secondmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
		//when middleware is executed, print out a message
		log.Print("Exectuing the second middleware")
		//if url path "other" is used then it will end here
		if r.URL.Path == "/other" {
			return
		}

		next.ServeHTTP(w,r) 
		//will print out a second message when it comes back
		log.Print("Executing second middleware again")
	})
}

//creating out handler function
func final(w http.ResponseWriter, r *http.Request) {
	log.Print("Executing the handler...")
	w.Write([]byte("Hello there!"))
}

func main() {
	//create our mutliplexer/router
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", firstmiddleware(secondmiddleware(finalHandler)))

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}