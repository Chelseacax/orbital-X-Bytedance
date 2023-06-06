package main	// package declaration 	
import (	// import declaration
	"fmt"	// import fmt package
	//"net/http"	// import net/http package
	"encoding/json"	// import encoding/json package
	"log"	// import log package
	//"io/ioutil"	// import ioutil package
)
//decode http request body into struct

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}


type Article struct {	// Article struct
	Id string `json:"Id"`	// Id string
	Title string `json:"Title"`	// Title string
	Desc string `json:"desc"`	// Desc string
	Content string `json:"content"`	// Content string
}

func main() {
    jsonUserString := `{"id": 1, "name": "John Doe", "email": "john.doe@example.com"}` // example
	jsonArticleString := `{"Id": "3", "Title": "Hello 3", "desc": "Article Description", "content": "Article Content"}`	// example
    
	var user User
	var article Article	// article variable

    err := json.Unmarshal([]byte(jsonUserString), &user)
    if err != nil {
        log.Fatal(err)
    }

	err = json.Unmarshal([]byte(jsonArticleString), &article)	// err variable
	if err != nil {		
		log.Fatal(err)	// log.Fatal function
	}

    fmt.Printf("User: %+v\n", user)
	fmt.Printf("Article: %+v\n", article)	// fmt.Printf function
}










// func homePage(w http.ResponseWriter, r *http.Request) {	// homePage function
// 	fmt.Fprintf(w, "Welcome to the HomePage!")	// fmt.Fprintf function
// 	fmt.Println("Endpoint Hit: homePage")	// fmt.Println function
// }

// func allArticles(w http.ResponseWriter, r *http.Request) { // allArticles function
// 	articles := Articles{	// articles variable
// 		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},	// Article struct
// 		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},	// Article struct
// 	}
// 	fmt.Println("Endpoint Hit: All Articles Endpoint")	// fmt.Println function
// 	json.NewEncoder(w).Encode(articles)	// json.NewEncoder function
// }

// func testPostArticles(w http.ResponseWriter, r *http.Request) {	// testPostArticles function
// 	fmt.Fprintf(w, "Test POST endpoint worked")	// fmt.Fprintf function
// 	fmt.Println("Endpoint Hit: testPostArticles")	// fmt.Println function
// }

// func returnSingleArticle(w http.ResponseWriter, r *http.Request) {	// returnSingleArticle function 
// 	vars := mux.Vars(r)	// vars variable
// 	key := vars["id"]	// key variable
// 	fmt.Fprintf(w, "Key: " + key)	// fmt.Fprintf function
// }

// func createNewArticle(w http.ResponseWriter, r *http.Request) {	// createNewArticle function 
// 	reqBody, _ := ioutil.ReadAll(r.Body)	// reqBody variable
// 	fmt.Fprintf(w, "%+v", string(reqBody))	// fmt.Fprintf function
// }

// func deleteArticle(w http.ResponseWriter, r *http.Request) {	// deleteArticle function
// 	fmt.Fprintf(w, "Endpoint Hit: deleteArticle")	// fmt.Fprintf function
// }

// func handleRequests() {	// handleRequests function
// 	myRouter := mux.NewRouter().StrictSlash(true)	// myRouter variable
// 	myRouter.HandleFunc("/", homePage)	// myRouter.HandleFunc function
// 	myRouter.HandleFunc("/articles", allArticles)	// myRouter.HandleFunc function
// 	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")	// myRouter.HandleFunc function
// 	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)	// myRouter.HandleFunc function
// 	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")	// myRouter.HandleFunc function
// 	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")	// myRouter.HandleFunc function
// 	log.Fatal(http.ListenAndServe(":8081", myRouter))	// log.Fatal function
// } 

// func main() {	// main function
// 	handleRequests()	// handleRequests function
// }

// // go run main.go
// // curl localhost:8081
// // curl localhost:8081/articles
// // curl -d '{"Id":"3","Title":"Hello 3","desc":"Article Description","content":"Article Content"}' -H "Content-Type: application/json" -X POST http://localhost:8081/articles
// // curl localhost:8081/articles/1
// // curl -d '{"Id":"3","Title":"Hello 3","desc":"Article Description","content":"Article Content"}' -H "Content-Type: application/json" -X POST http://localhost:8081/articles
// // curl -X DELETE localhost:8081/articles/1

// // Endpoint Hit: homePage
// // Endpoint Hit: All Articles Endpoint
// // Endpoint Hit: testPostArticles
// // Endpoint Hit: deleteArticle

