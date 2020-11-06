package main

import (
	"fmt"
	"net/http"
)

func main(){	
	http.HandleFunc("/", useCPU)
	http.ListenAndServe(":80", nil)
}

func useCPU(w http.ResponseWriter, r *http.Request) {
	count := 1

	for i := 1; i <= 1000000; i++ {
		count = i
	}

	fmt.Printf("count: %d", count)
	w.Write([]byte(string(count)))
}
