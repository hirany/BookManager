package main

import (
	"github.com/antonholmquist/jason"
	"net/http"
	"strconv"
)

func janToTitle(bookID int64) (string, error) {
	url := "https://www.googleapis.com/books/v1/volumes?q="
	resp, err := http.Get(url + strconv.FormatInt(bookID, 10))
	if err != nil {
		return "none", err
	}
	defer resp.Body.Close()
	v, err := jason.NewObjectFromReader(resp.Body)
	if err != nil {
		return "none", err
	}
	items, err := v.GetObjectArray("items")
	if err != nil {
		return "none", err
	}
	title, err := items[0].GetString("volumeInfo", "title")
	if err != nil {
		return "none", err
	}
	return title, nil
}
