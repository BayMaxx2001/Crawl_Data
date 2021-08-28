package page

import (
	"bufio"
	"log"
	"strings"

	"crawl_data/model"
	"crawl_data/utils"
)

func GetInformationOfDate(dayURL string) []model.PageInformation {
	var (
		listInformationDate []model.PageInformation
		informationDate     model.PageInformation
		md5                 string
		sha1                string
		sha256              string
		day                 string
		month               string
		year                string
	)

	year, month, day = utils.GetDateDetail(dayURL)
	informationDate = model.NewPageInformation("", "", "", day, month, year)

	log.Println("Day", dayURL, ": Request get information ")
	urlDate := utils.GetAddressOfDate(dayURL)
	resp := utils.RequestHTTP(urlDate)

	if resp == nil {
		listInformationDate = append(listInformationDate, informationDate)
		return listInformationDate
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("Status code error: %d, %s", resp.StatusCode, resp.Status)
		listInformationDate = append(listInformationDate, informationDate)
		return listInformationDate
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		words := strings.Fields(scanner.Text())
		if len(words) < 3 {
			continue
		}
		md5 = words[0]
		sha1 = words[1]
		sha256 = words[2]
		informationDate = model.NewPageInformation(md5, sha1, sha256, day, month, year)
		listInformationDate = append(listInformationDate, informationDate)
	}
	if err := scanner.Err(); err != nil {
		log.Println("Error get date information ", err)
	}

	if len(listInformationDate) == 0 {
		listInformationDate = append(listInformationDate, informationDate)
	}
	log.Println("Get information", dayURL, ": Successfully")
	return listInformationDate

}
