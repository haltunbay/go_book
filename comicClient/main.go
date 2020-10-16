package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
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

var comicTemplate = `
<table>
<tr>
	<td>#</td><td>Month/Year</td><td>Title</td><td>Transcript</td>
</tr>
{{range .}}
	<tr>
		<td>{{.Num}}</td><td>{{.Month}}-{{.Year}}</td><td>{{.Title}}</td><td>{{.Transcript}}</td>
	</tr>
{{end}}
</table>
`
var report = template.Must(template.New("comics").Parse(comicTemplate))

func main() {
	buildFromTemplate()
	buildIndex()
	term := os.Args[1]
	m, ok := index[term]
	if !ok {
		fmt.Printf("term: %s not found\n", term)
	} else {
		fmt.Println("term found in following chapters", m)
	}
}

func buildFromTemplate() {
	comics := fetchComics()
	f, err := os.Create("out.html")
	check(err)
	report.Execute(f, comics)
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

func fetchComics() []Comic {
	resp, err := http.Get("http://localhost:8080/comics")
	check(err)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	var comics []Comic
	json.Unmarshal(data, &comics)
	return comics
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
