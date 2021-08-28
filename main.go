package main

import (
	"fmt"
	"sync"
	"time"

	"crawl_data/model"
	"crawl_data/page"
	"crawl_data/utils"
)

var url string = "https://malshare.com/daily/"

func sendDateToChanDate(chanDate chan<- string, listDate []string) {
	for i := range listDate {
		chanDate <- listDate[i]
	}
	close(chanDate)
}

func processInformationOfDate(chanDate chan string, chanPageInformation chan []model.PageInformation, wg *sync.WaitGroup) {
	for date := range chanDate {
		chanPageInformation <- page.GetInformationOfDate(date)
	}
	wg.Done()
}

func spawnWorkerDate(numOfWorkerDate int, chanDate chan string, chanPageInformation chan []model.PageInformation) {
	var wg sync.WaitGroup
	for i := 0; i < numOfWorkerDate; i++ {
		wg.Add(1)
		go processInformationOfDate(chanDate, chanPageInformation, &wg)
	}
	wg.Wait()
	close(chanPageInformation)
}

func writeOutputToFile(chanPageInformation chan []model.PageInformation, wg *sync.WaitGroup) {
	for info := range chanPageInformation {
		day := info[0].DAY
		month := info[0].MONTH
		year := info[0].YEAR

		date := utils.GetDateToString(day, month, year)
		go utils.SaveFile(date, "md5", info)
		go utils.SaveFile(date, "sha1", info)
		go utils.SaveFile(date, "sha256", info)
	}
	wg.Done()
}

func spawnWorkerWriteOutput(done chan bool, numOfWorkerPageInformation int, chanPageInformation chan []model.PageInformation) {
	var wg sync.WaitGroup
	for i := 0; i < numOfWorkerPageInformation; i++ {
		wg.Add(1)
		go writeOutputToFile(chanPageInformation, &wg)
	}
	wg.Wait()
	done <- true
}

func crawlData() {
	var (
		done                       = make(chan bool)
		chanPageInformation        = make(chan []model.PageInformation, 500)
		chanDate                   = make(chan string, 500)
		listDate                   []string
		numOfWorkerDate            = 100
		numOfWorkerPageInformation = 100
	)

	listDate = utils.GetListDate(url)
	go sendDateToChanDate(chanDate, listDate)
	go spawnWorkerWriteOutput(done, numOfWorkerPageInformation, chanPageInformation)
	spawnWorkerDate(numOfWorkerDate, chanDate, chanPageInformation)
	<-done
}

func main() {
	start := time.Now()
	crawlData()
	fmt.Println(time.Since(start))
}
