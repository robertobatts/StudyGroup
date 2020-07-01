package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type WebComic struct {
	Alt        string `json:"alt"`
	Day        string `json:"day"`
	Img        string `json:"img"`
	Link       string `json:"link"`
	Month      string `json:"month"`
	News       string `json:"news"`
	Num        int    `json:"num"`
	SafeTitle  string `json:"safe_title"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Year       string `json:"year"`
}

const webComicURL = "http://xkcd.com/"
const infoURLSuffix = "/info.0.json"

var comics = make(map[string][]WebComic)

func run() {
	downloadComics(1, 10)
	lookup(os.Args[1:])
}

func lookup(words []string) {
	results := make(map[WebComic]string)
	for _, word := range words {
		wResults := comics[word]
		for _, wRes := range wResults {
			results[wRes] += word + " "
		}
	}
	for comic, match := range results {
		fmt.Println("matched by " + match + ":")
		fmt.Println(comic.Transcript)
		fmt.Println(comic.Img + "\n")
	}
}

func downloadComics(from, to int) {
	for i := from; i <= to; i++ {
		url := webComicURL + strconv.Itoa(i) + infoURLSuffix
		downloadAndIndex(url)
	}
}

func downloadAndIndex(url string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	var result WebComic
	json.NewDecoder(resp.Body).Decode(&result)

	transcript := result.Transcript
	removeSpecialChars(&transcript)
	words := strings.Split(transcript, " ")
	for _, word := range words {
		comicSlice := comics[word]
		comicSlice = append(comicSlice, result)
		comics[word] = comicSlice
	}
}

func removeSpecialChars(s *string) {
	*s = strings.Replace(*s, "[", " ", -1)
	*s = strings.Replace(*s, "]", " ", -1)
	*s = strings.Replace(*s, ".", " ", -1)
	*s = strings.Replace(*s, ":", " ", -1)
	*s = strings.Replace(*s, "\\n", " ", -1)
}
