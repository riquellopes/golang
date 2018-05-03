package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RSS -
type RSS struct {
	XMLName     xml.Name `xml:"rss"`
	Version     string   `xml:"version,attr"`
	Title       string   `xml:"channel>title"`
	Link        string   `xml:"channel>link"`
	Description string   `xml:"channel>description"`
	PubDate     string   `xml:"channel>pubDate"`
	ItemList    []Item   `xml:"channel>item"`
}

// Item -
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Content     string `xml:"encoded"`
	PubDate     string `xml:"pubDate"`
	Comments    string `xml:"comments"`
}

func main() {
	resp, err := http.Get("http://pizzadedados.com/feed.xml")

	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	rss := RSS{}
	err = xml.Unmarshal(body, &rss)

	if err != nil {
		fmt.Println(err)
	}

	for k, i := range rss.ItemList {
		fmt.Println(k, i.Title)
	}
}
