package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func normalizeURL(str string) (string, error) {
	url, err := url.Parse(str)
	actual := url.Scheme + "://" + url.Host + url.Path

	return actual, err

}

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal("error create newdocumentfromreader")
	}
	h1 := doc.Find("h1").First().Text()
	return h1

}
func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal("error create newdocumentfromreader")
	}
	var p string
	main := doc.Find("main")
	if main.Size() > 0 {
		p = main.Find("p").First().Text()
		return p
	}
	p = doc.Find("p").First().Text()
	return p

}

func getURLsFromHTML(body string, baseUrl *url.URL) ([]string, error) {
	actual := []string{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Fatal("error create newdocumentfromreader")
	}
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		val, ok := s.Attr("href")
		if !ok {
			log.Fatal("not found")
		}
		idx := strings.Index(val, "/")
		if idx == 0 {
			val = baseUrl.Scheme + "://" + baseUrl.Host + baseUrl.Path + val[idx+1:]
		}
		actual = append(actual, val)

	})

	return actual, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	actual := []string{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		log.Fatal("error create newdocumentfromreader")
	}
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		val, ok := s.Attr("src")
		if !ok {
			log.Fatal("Not found")
		}
		idx := strings.Index(val, "/")

		if idx == 0 {
			val = baseURL.Scheme + "://" + baseURL.Host + baseURL.Path + val
		}
		actual = append(actual, val)
	})

	return actual, nil
}

type PageData struct {
	URL             string
	H1              string
	First_paragraph string
	Outgoing_links  []string
	Image_urls      []string
	Count           int
}

func extractPageData(html, pageURL string) PageData {
	pageData := PageData{}

	url, err := url.Parse(pageURL)
	if err != nil {
		log.Fatal("Error parsing url")
	}

	pageData.Outgoing_links, err = getURLsFromHTML(html, url)
	if err != nil {
		log.Fatal("Error extracting url")
	}
	pageData.H1 = getH1FromHTML(html)
	pageData.First_paragraph = getFirstParagraphFromHTML(html)
	pageData.URL, err = normalizeURL(pageURL)
	if err != nil {
		log.Fatal("Error normalize URL")
	}
	pageData.Image_urls, err = getImagesFromHTML(html, url)
	if err != nil {
		log.Fatal("Error extracting image url")
	}

	return pageData
}
