package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

type Config struct {
	pages              map[string]PageData
	baseUrl            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	args := os.Args
	cfg := Config{pages: map[string]PageData{}}
	cfg.crawlPage(args[1])
	fmt.Println("Finished crawl")

}

func (cfg *Config) crawlPage(rawCurrentUrl string) error {
	body, err := getHTML(rawCurrentUrl)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	cfg.pages[rawCurrentUrl] = extractPageData(body, rawCurrentUrl)
	for _, link := range cfg.pages[rawCurrentUrl].Outgoing_links {
		if strings.HasPrefix(link, rawCurrentUrl) {
			if _, visited := cfg.pages[link]; !visited {
				fmt.Println("Start crawling at ", link)
				cfg.crawlPage(link)
			}
		}
	}

	return nil
}

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		panic(fmt.Sprintf("Error making new request: %v", err))
	}
	req.Header.Set("User-Agent", "BootCrawler/1.0")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error getting response: %v", err)
	}
	if resp.StatusCode > 400 {
		return "", errors.New("Responding with 400+ status code")
	}
	// if contentType := resp.Header.Get("Content-Type"); contentType != "text/html" {
	// 	return "", errors.New("The site doesn't contain valid html")
	// }
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		panic(fmt.Sprintf("Error reading the body: %v", err))
	}
	return string(body), nil
}
