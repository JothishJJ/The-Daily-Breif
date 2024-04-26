package router

import (
	"net/http"

	"github.com/JothishJJ/thedailybrief/crawler"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	r.GET("/world", func(c *gin.Context) {

		crawlerJobs := crawler.Websites{
			"https://www.thehindu.com/news/international/feeder/default.rss",
			"https://timesofindia.indiatimes.com/rssfeeds/296589292.cms",
		}

		allDatas := crawler.GetAllData(crawlerJobs)

		c.JSON(http.StatusOK, allDatas)
	})
}
