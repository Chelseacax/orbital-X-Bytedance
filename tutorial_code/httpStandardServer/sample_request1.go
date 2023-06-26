package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func main() {
	// The http.Get function makes a GET request to the specified URL.
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var post Post
	json.Unmarshal(body, &post)

	fmt.Printf("Post ID: %s\n", post.ID)
	fmt.Printf("Post Title: %s\n", post.Title)
}
