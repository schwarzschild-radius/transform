package main

import (
	"net/http"
	"./spmdfy"
	"log"
	"bytes"
)

func handler(response http.ResponseWriter, request *http.Request) {
	log.SetPrefix("[transform/spmdfy_handler] ")
	src := new(bytes.Buffer)
	src.ReadFrom(request.Body)
	log.Println("initiating compilation...")
	out, _ := spmdfy.Spmdfy(src.String())
	response.Write([]byte(out))
}

func main() {
	log.SetPrefix("[transform/main] ")
	log.Println("Starting http server at port" + "8000")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
