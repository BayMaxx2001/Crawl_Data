package utils

import (
	"log"
)

func GetListDate(url string) []string {
	var (
		listDate []string
	)
	log.Println("Start getting the list date")

	resp := RequestHTTP(url)
	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d, %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc := LoadHTML(resp)

	// Find the day items
	sel_getTr := doc.Find("table tr")
	for i := range sel_getTr.Nodes {

		query := sel_getTr.Eq(i)
		sel_getTd := query.Find("td")

		for j := range sel_getTd.Nodes {

			day := sel_getTd.Eq(j)

			if j == 1 {
				add := string(day.Text())
				if IsDate(add) == true {
					listDate = append(listDate, string(add[:len(add)-1]))
				}
				//log.Println(string(day.Text()))
			}
		}
	}

	log.Println("Complete get the list date")
	return listDate
}
