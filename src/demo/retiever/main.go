package main

import (
	"fmt"
	"GoTest/src/demo/retiever/mock"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "allen",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster)string {
	s.Post(url,map[string]string{
		"contents":"another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	s:=mock.Retriever{"abcd"}
	fmt.Println(session(&s))
}
