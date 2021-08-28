package utils

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func RequestHTTP(url string) *http.Response {

	// Request http
	log.Println("Request http get to", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error request to url: ", err)
		return nil
	}
	return resp
}

func LoadHTML(resp *http.Response) *goquery.Document {
	for i := 1; i <= 5; i++ {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatalln("Error load html: ", err)
		} else {
			return doc
		}
	}
	return nil
}
