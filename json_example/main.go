package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

const api string = "https://hacker-news.firebaseio.com/v0/item/8863.json?print=pretty"

type item struct {
	By          string `json:"by"`
	Descendants int32  `json:"descendants"`
	Id          int32  `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int32  `json:"score"`
	Time        int32  `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Url         string `json:"url"`
}

func main() {

	resp, err := http.Get(api)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	defer resp.Body.Close()

	var itemResponse item
	err = json.NewDecoder(resp.Body).Decode(&itemResponse)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(itemResponse.Type)
}
