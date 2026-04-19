package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MetObject struct {
	Title             string `json:"title"`
	ArtistDisplayName string `json:"artistDisplayName"`
	ObjectDate        string `json:"objectDate"`
	PrimaryImage      string `json:"primaryImage"`
}

func main() {
	fmt.Println("Hello World")

	objectID := "436532"
	apiURL := fmt.Sprintf("https://collectionapi.metmuseum.org/public/collection/v1/objects/%s", objectID)

	resp, err := http.Get(apiURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var obj MetObject
	err = json.NewDecoder(resp.Body).Decode(&obj)
	if err != nil {
		panic(err)
	}

	fmt.Println(obj.Title)
	fmt.Println(obj.PrimaryImage)
	http.Get(obj.PrimaryImage)
}
