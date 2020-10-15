package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var comics []Comic

const ComicUrl = "http://localhost:8080/%d/info.0.json"

type Comic struct {
	Month      int    `json:"month"`
	Year       int    `json:"year"`
	Num        int    `json:"num"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {
	buildIndex()
}

func buildIndex() {
	for i := 1; i < 5; i++ {
		url := fmt.Sprintf(ComicUrl, i)
		//fmt.Println("fetching url", url)

		c := fetchComic(url)
		fmt.Println(c.Transcript)
		for _, t, err := bufio.ScanWords([]byte(c.Transcript)) {
			if err != nil{
				fmt.Errorf("Error scanwords %s", err)
			}
		}
	}
}

func fetchComic(url string) Comic {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	var comic Comic
	json.Unmarshal(data, &comic)
	return comic
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
