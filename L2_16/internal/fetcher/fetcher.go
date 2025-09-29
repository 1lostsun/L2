package fetcher

import (
	"errors"
	Io "io"
	"log"
	"net/http"
	"time"
)

// Fetcher : Структура отправки HTTP-запросов
type Fetcher struct {
	client *http.Client
}

// NewFetcher : Конструктор структуры фетчера, по умолчанию 30 секунд таймаут
func NewFetcher() *Fetcher {
	return &Fetcher{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Fetch : загружает содержимое страницы по указанному URL.
func (f *Fetcher) Fetch(url string) ([]byte, error) {
	resp, err := f.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer func(Body Io.ReadCloser) {
		closeErr := Body.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	content, err := Io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
