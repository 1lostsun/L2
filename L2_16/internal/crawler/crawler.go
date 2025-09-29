package crawler

import (
	"fmt"
	"github.com/1lostsun/L2/tree/main/L2_16/internal/fetcher"
	"github.com/1lostsun/L2/tree/main/L2_16/internal/parser"
	"github.com/1lostsun/L2/tree/main/L2_16/internal/storage"
)

// Crawler : Структура кравлера веб страниц
type Crawler struct {
	*fetcher.Fetcher
	*storage.Storage
	*parser.Parser
	visited  map[string]bool
	maxDepth int
}

// NewCrawler : Конструктор кравлера
func NewCrawler(folderName string, maxDepth int) *Crawler {
	return &Crawler{
		Fetcher:  fetcher.NewFetcher(),
		Storage:  storage.NewStorageOnDesktop(folderName),
		Parser:   parser.NewParser(),
		visited:  make(map[string]bool),
		maxDepth: maxDepth,
	}
}

// Crawl : Вызов функции краулинга, начиная с переданного URL
func (c *Crawler) Crawl(url string) error {
	return c.crawlRecursively(url, 0)
}

func (c *Crawler) crawlRecursively(URL string, depth int) error {
	if depth > c.maxDepth {
		fmt.Printf("max depth reached %d for %s\n", c.maxDepth, URL)
		return nil
	}

	if c.visited[URL] {
		fmt.Printf("url: %s\n already visited ", URL)
		return nil
	}

	c.visited[URL] = true

	content, fetchErr := c.Fetch(URL) // делаем http запрос
	if fetchErr != nil {
		return fetchErr
	}

	if storeErr := c.Save(URL, content); storeErr != nil { // сохраняем в файл на рабочем столе
		return storeErr
	}

	links, extractErr := c.ExtractLinks(string(content), URL)
	if extractErr != nil {
		return extractErr
	}

	for i, link := range links {
		if i >= 20 {
			fmt.Printf("only 20 links on page")
			break
		}

		if c.visited[link] {
			continue
		}

		if err := c.crawlRecursively(link, depth+1); err != nil {
			fmt.Printf("crawl recursively error: %v\n", err)
		}
	}

	return nil
}
