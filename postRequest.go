package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	// We're using `strings.NewReader` to create an `io.Reader` with our request body.
	reader := strings.NewReader(`{"title":"foo","body":"bar","userId":1}`)
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", reader)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))
}

