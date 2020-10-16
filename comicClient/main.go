package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var index = map[string]map[int]bool{}

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
	term := os.Args[1]
	fmt.Println(term)
	m, ok := index[term]
	if !ok {
		fmt.Printf("term: %s nor found\n", term)
	} else {
		fmt.Println("term found in following chapters", m)
	}
}

func buildIndex() {
	for i := 1; i < 5; i++ {
		url := fmt.Sprintf(ComicUrl, i)
		//fmt.Println("fetching url", url)

		c := fetchComic(url)
		//fmt.Println(c.Transcript)
		scanner := bufio.NewScanner(strings.NewReader(c.Transcript))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			w := cleanup(scanner.Text())

			if index[w] == nil {
				index[w] = map[int]bool{}
			}
			wordmap := index[w]

			wordmap[i] = true
		}
	}
	//data, _ := json.MarshalIndent(index, "", " ")
	//fmt.Printf("%+s \n", string(data))

}

func cleanup(s string) string {
	list := []string{"[[", "]]", "{{", "}}"}
	for _, sp := range list {
		s = strings.Trim(s, sp)
	}
	return s
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
