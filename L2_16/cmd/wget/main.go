package main

import "C"
import (
	"fmt"
	"github.com/1lostsun/L2/tree/main/L2_16/internal/crawler"
	"github.com/1lostsun/L2/tree/main/L2_16/internal/logger"
	"log"
)

func main() {
	Logger := logger.New()
	var url string
	Logger.Log(false, "input url for downloading: ")

	_, scanErr := fmt.Scanln(&url)
	if scanErr != nil {
		log.Fatal(scanErr)
	}

	var maxDepth int
	Logger.Log(false, "input the max crawl depth: ")
	_, scanErr = fmt.Scanln(&maxDepth)
	if scanErr != nil {
		log.Fatal(scanErr)
	}

	if maxDepth <= 0 {
		maxDepth = 2
	}

	cr := Crawler.NewCrawler("crawler_downloads", maxDepth)
	Logger.Log(true, "files will downloads in: ", cr.GetBaseDir())

	err := cr.Crawl(url)
	if err != nil {
		log.Fatal(err)
	}

	Logger.Log(true, "crawling was ended")
	Logger.Log(true, "files was downloads in: ", cr.GetBaseDir())
}
