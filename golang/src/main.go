package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io"
	"io/ioutil"
)

var configuration []byte
var secret []byte

func Response(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from Go")
}

func Status(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func ReadConfig(){
	fmt.Println("reading config...")
	config, e := ioutil.ReadFile("/configs/config.json")
	if e != nil {
		fmt.Printf("Error reading config file: %v\n", e)
		os.Exit(1)
	}
	configuration = config
	fmt.Println("config loaded!")

}

func ReadSecret(){
	fmt.Println("reading secret...")
	s, e := ioutil.ReadFile("/secrets/secret.json")
	if e != nil {
		fmt.Printf("Error reading secret file: %v\n", e)
		os.Exit(1)
	}
	secret = s
	fmt.Println("secret loaded!")

}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log basic request information
		log.Printf(
			"%s %s %s %s", // Format: [Remote Address] [Method] [Path] [Protocol]
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			r.Proto,
		)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func main() {
    
	fmt.Println("starting...")
	ReadConfig()
	ReadSecret()
	
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", Response)
	mux.HandleFunc("GET /status", Status)
	loggedMux := Logger(mux)

	server := &http.Server{
		Addr:    ":5000",
		Handler: loggedMux,
	}

	log.Fatal(server.ListenAndServe())
}
