package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/opesun/goquery"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func fileNameParse(site string) string {
	urls := strings.Split(site, "/")
	return urls[2] + ".html"
}

func download(site string) {
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fileName := fileNameParse(site)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
}

func parseResources(site string) {
	x, _ := goquery.ParseUrl(site)
	for _, url := range x.Find("").Attrs("href") {
		var str []string
		switch {
		case strings.Contains(url, ".png"):
			str = strings.Split(url, "/")
			downloadResources(str[len(str)-1], url)
		case strings.Contains(url, ".jpg"):
			str = strings.Split(url, "/")
			downloadResources(str[len(str)-1], url)
		case strings.Contains(url, ".css"):
			str = strings.Split(url, "/")
			downloadResources(str[len(str)-1], url)
		}
	}
}

func downloadResources(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func Wget() {
	var site string
	flag.StringVar(&site, "s", "https://www.google.com/", "site")

	flag.Parse()

	if ok, err := regexp.MatchString("^(http|https)://", site); ok == true && err == nil {
		download(site)
		parseResources(site)
	} else {
		log.Fatal("invalid url")
	}
}

func main() {
	Wget()
}
