package main

import (
	"net/http"

	"GoTest/src/demo/crawler/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("src/demo/crawler/frontend/view")))
	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"src/demo/crawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
