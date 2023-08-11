package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/mmcdole/gofeed/rss"
)

func main() {
	files, _ := os.ReadDir("./")
	for _, file := range files {
		if strings.Contains(file.Name(), ".mp3") {
			os.Remove(file.Name())
		}
	}
	resp, err := http.Get("https://feeds.simplecast.com/54nAGcIl")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fp := rss.Parser{}
	rssFeed, _ := fp.Parse(strings.NewReader(string(body)))
	resp, err = http.Get(rssFeed.Items[0].Enclosure.URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(rssFeed.Items[0].PubDateParsed.Local().Format("2006/01/02")))
	filename := rssFeed.Items[0].Title + ".mp3"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(file, resp.Body)
}
