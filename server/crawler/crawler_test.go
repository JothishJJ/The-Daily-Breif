package crawler_test

import (
	"testing"

	"github.com/JothishJJ/thedailybrief/crawler"
)

func TestCrawler(t *testing.T) {
	_, err := crawler.Crawler("https://timesofindia.indiatimes.com/rssfeeds/296589292.cms")
	if err != nil {
		t.Fatalf("Crawler Test failed: %v", err)
	}
}

func TestGetAllData(t *testing.T) {
	websites := crawler.Websites{
		"https://timesofindia.indiatimes.com/rssfeeds/296589292.cms",
	}

	datas := crawler.GetAllData(websites)

	for _, data := range datas {
		if len(data) > 20 {
			t.Fatal("Data has more than 20 items")
		}
	}
}
