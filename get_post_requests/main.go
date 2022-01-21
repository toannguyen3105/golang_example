package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func getData() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Println(sb)
}

func main() {
	//getData()

	postBody, _ := json.Marshal(map[string]string{
		"name":  "Toby",
		"email": "Toby@example.com",
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
