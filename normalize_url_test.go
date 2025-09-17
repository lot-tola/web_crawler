package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><a href="https://blog.boot.dev"><span>Boot.dev</span></a></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
func TestGetURLsFromHTMLRelative(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><a href="/catalogue/index.html"><span>Boot.dev</span></a></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getURLsFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/catalogue/index.html"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
func TestGetImagesFromHTMLRelative(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><img src="/logo.png" alt="Logo"></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetImagesFromHTMLAbsolute(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><img src="https://blog.boot.dev/logo.png" alt="Logo"></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/logo.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestGetImagesFromHTMLAbsolute1(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body><img src="https://blog.boot.dev/lane.png" alt="Logo"><img src="https://blog.boot.dev/picture.png" alt="Logo"></body></html>`

	baseURL, err := url.Parse(inputURL)
	if err != nil {
		t.Errorf("couldn't parse input URL: %v", err)
		return
	}

	actual, err := getImagesFromHTML(inputBody, baseURL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"https://blog.boot.dev/lane.png", "https://blog.boot.dev/picture.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestExtractPageDataBasic(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body>
        <h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
    </body></html>`

	actual := extractPageData(inputBody, inputURL)

	expected := PageData{
		URL:             "https://blog.boot.dev",
		H1:              "Test Title",
		First_paragraph: "This is the first paragraph.",
		Outgoing_links:  []string{"https://blog.boot.dev/link1"},
		Image_urls:      []string{"https://blog.boot.dev/image1.jpg"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %#v, got %#v", expected, actual)
	}
}
