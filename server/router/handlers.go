package router

import (
	"net/http"

	"github.com/JothishJJ/thedailybrief/crawler"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func TopHandler(c *gin.Context) {
	crawlerJobs := crawler.Websites{
		"https://timesofindia.indiatimes.com/rssfeedstopstories.cms",
	}

	allDatas := crawler.GetAllData(crawlerJobs)

	c.JSON(http.StatusOK, allDatas)
}

func WorldHandler(c *gin.Context) {

	crawlerJobs := crawler.Websites{
		"https://timesofindia.indiatimes.com/rssfeeds/296589292.cms",
	}

	allDatas := crawler.GetAllData(crawlerJobs)

	c.JSON(http.StatusOK, allDatas)
}

func IndiaHandler(c *gin.Context) {
	crawlerJobs := crawler.Websites{
		"https://timesofindia.indiatimes.com/rssfeeds/-2128936835.cms",
	}

	allDatas := crawler.GetAllData(crawlerJobs)

	c.JSON(http.StatusOK, allDatas)
}
