package parser

import (
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

// Parser : Структура парсера
type Parser struct{}

// NewParser : Конструктор парсера
func NewParser() *Parser {
	return &Parser{}
}

// ExtractLinks : Выделяет все ссылки из HTML-документа, учитывая baseURL.
func (p *Parser) ExtractLinks(htmlContent, baseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var links []string
	p.findLinks(doc, baseURL, &links)
	return links, nil
}

func (p *Parser) findLinks(n *html.Node, baseURL string, links *[]string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if absolute := toAbsoluteURL(a.Val, baseURL); absolute != "" {
					*links = append(*links, a.Val)
				}
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		p.findLinks(c, baseURL, links)
	}
}

func toAbsoluteURL(href, baseURL string) string {
	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}

	link, err := url.Parse(href)
	if err != nil {
		return ""
	}

	absolute := base.ResolveReference(link)
	if absolute.Scheme == "http" || absolute.Scheme == "https" {
		return absolute.String()
	}

	return ""
}
