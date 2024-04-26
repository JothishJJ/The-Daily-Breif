package crawler

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Item struct {
	Title     string    `xml:"title"`
	Link      string    `xml:"link"`
	Desc      string    `xml:"description"`
	City      string    `xml:"city"`
	Logo      string    `xml:"logo"`
	JobType   string    `xml:"jobtype"`
	Category  string    `xml:"category"`
	PubDate   string    `xml:"date"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
	Items []Item `xml:"item"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Websites = []string

type datas = [][]Item

/*
Crawls a website or RSS and gets the items inside it.
It returns either an []Item and an error
*/
func Crawler(address string) ([]Item, error) {

	rss := Rss{}

	resp, err := http.Get(address)
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return rss.Channel.Items, err
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
		return rss.Channel.Items, err
	}
	var data []Item
	for _, item := range rss.Channel.Items {
		data = append(data, item)
	}

	return data, nil
}

/*
Accepts a []string and then use the crawl function to encode all of them and returns as an array.
It accepts a []string and return a [][]Item
*/
func GetAllData(crawlerJobs Websites) datas {
	allDatas := datas{}

	for _, crawlerJob := range crawlerJobs {
		data, err := Crawler(crawlerJob)
		if err != nil {
			log.Fatal(err)
		}

		if len(data) > 20 {
			allDatas = append(allDatas, data[:20])
		} else {
			allDatas = append(allDatas, data)
		}
	}

	return allDatas
}
