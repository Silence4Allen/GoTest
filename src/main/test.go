package main

import (
	"strconv"
	"fmt"
)

const SongListUrlModel = "http://www.app-echo.com/api/channel/info?order=hot&id=1155&page="

func main() {
	for i := 1; i < 40; i++ {
		url := SongListUrlModel + strconv.Itoa(i)
		fmt.Println(url)
	}
}
