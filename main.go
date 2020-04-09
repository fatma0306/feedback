package main

import (
	"fmt"
	"net/http"

	"github.com/fatma0306/feedback/handler"
)

func main() {
	r := handler.Router()
	fmt.Println("Starting server on the port 5000...")
	err := http.ListenAndServe(":5000", r)
	if err != nil {
		fmt.Println(err)
	}
}
