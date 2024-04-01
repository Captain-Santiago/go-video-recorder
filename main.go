package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

//go:embed front
var front embed.FS

func main_page(w http.ResponseWriter, req *http.Request) {
	html, err := front.ReadFile("front/index.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(html))
}

func start_recording(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Start")
}

func stop_recording(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Stop recording")
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	http.HandleFunc("/", main_page)
	http.HandleFunc("/start", start_recording)
	http.HandleFunc("/stop", stop_recording)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", SERVER_PORT), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Printf("Server running: http://localhost:%d\nCTRL+C to stop\n", SERVER_PORT)

	<-c
}
