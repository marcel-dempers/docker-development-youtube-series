package main

import (
    "net/http"
    "log"
		"encoding/json"
)

type WeatherData struct {
	Location struct {
			City    string `json:"city"`
			Region  string `json:"region"`
			Country string `json:"country"`
	} `json:"location"`
	Current struct {
			Temperature   int    `json:"temperature"`
			Humidity      int    `json:"humidity"`
			WindSpeed     int    `json:"wind_speed"`
			WindDirection string `json:"wind_direction"`
			Condition     string `json:"condition"`
	} `json:"current"`
	Forecast []struct {
			Date      string `json:"date"`
			High      int    `json:"high"`
			Low       int    `json:"low"`
			Condition string `json:"condition"`
	} `json:"forecast"`
}

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			
			sw := &statusResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(sw, r)
			log.Printf("%s %s %s %s %d\n", r.Method, r.URL.Path, r.RemoteAddr, r.Proto, sw.statusCode)
	})
}


func weatherHandler(w http.ResponseWriter, r *http.Request) {
	weather := WeatherData{
			Location: struct {
					City    string `json:"city"`
					Region  string `json:"region"`
					Country string `json:"country"`
			}{
					City:    "San Francisco",
					Region:  "CA",
					Country: "USA",
			},
			Current: struct {
					Temperature   int    `json:"temperature"`
					Humidity      int    `json:"humidity"`
					WindSpeed     int    `json:"wind_speed"`
					WindDirection string `json:"wind_direction"`
					Condition     string `json:"condition"`
			}{
					Temperature:   18,
					Humidity:      72,
					WindSpeed:     12,
					WindDirection: "NW",
					Condition:     "Partly Cloudy",
			},
			Forecast: []struct {
					Date      string `json:"date"`
					High      int    `json:"high"`
					Low       int    `json:"low"`
					Condition string `json:"condition"`
			}{
					{Date: "2025-04-24", High: 20, Low: 14, Condition: "Sunny"},
					{Date: "2025-04-25", High: 19, Low: 13, Condition: "Cloudy"},
					{Date: "2025-04-26", High: 21, Low: 15, Condition: "Rain"},
			},
	}

  sw := &statusResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
	log.Printf("%s %s %s %s %d\n", r.Method, r.URL.Path, r.RemoteAddr, r.Proto, sw.statusCode)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(weather); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func main() {
    
		fileServer := http.FileServer(http.Dir("."))
		
		http.Handle("/", loggingMiddleware(fileServer))
    http.HandleFunc("/weather", weatherHandler)

		log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}