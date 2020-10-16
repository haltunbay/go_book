package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Comic struct {
	Month      int    `json:"month"`
	Year       int    `json:"year"`
	Num        int    `json:"num"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

var comics []Comic

func main() {
	http.HandleFunc("/", handler)
	loadComics()
	//fmt.Printf("%+v", comics)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	ci := path[1]
	index, _ := strconv.Atoi(ci)
	json, err := json.MarshalIndent(comics[index-1], "", " ")
	check(err)

	fmt.Fprint(w, string(json))
}

func loadComics() {
	file, err := ioutil.ReadFile("data.json")
	check(err)
	json.Unmarshal(file, &comics)
	//fmt.Printf("%+v", comics)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
