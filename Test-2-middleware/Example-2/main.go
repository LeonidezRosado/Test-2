package main

import (
	"log"
	"net/http"
	"mime"
)

//creating out json authentication middleware
func authenticationhandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		//if content type is a defect then display a message
		if contentType != "" {
			mt, _,err := mime.ParseMediaType(contentType) 
			if err != nil {
				http.Error(w, "defect content-Type header", http.StatusBadRequest)
				return
			}

			//if content-Type is not of application/json then display a message
			if mt != "application/json" {
					http.Error(w, "The Content-Type must be of application/json", http.StatusUnsupportedMediaType)
					return
			}
		}
		next.ServeHTTP(w,r)
	})
}

//create our handler function
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	//create our multiplexer/router
	mux := http.NewServeMux()
	mux.Handle("/", authenticationhandler(http.HandlerFunc(handler)))

	//creating our server on our port
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}