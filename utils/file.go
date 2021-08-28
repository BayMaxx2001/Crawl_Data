package utils

import (
	"log"
	"os"

	"crawl_data/model"
)

func NameFile(day string, format string) string {
	nameFile := string(day[:len(day)-1]) + "." + format + ".txt"
	return nameFile
}

// Create folder contains information of the date
func CreateFolderDay(day string, month string, year string) string {

	path := "output/" + year + "/" + month + "/" + day + "/"
	err := os.MkdirAll(path, 0775)
	if err != nil {
		log.Println("Error Create Folder: ", year+"/"+month+"/"+day, err)
	}
	return path
}

func WriteFile(path string, content string, day string) {
	log.Println("Day", day, "write information to file")

	f, err := os.Create(path)
	if err != nil {
		log.Println("Cannot create folder date:", err)
		return
	}

	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Cannot open file", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(content + "\n"); err != nil {
		log.Println("Cannot write", err)
		f.Close()
	}

	err = f.Close()
	if err != nil {
		log.Println(err)
	}
	log.Println("Write information successfully")
}

func SaveFile(date string, format string, content []model.PageInformation) {

	year, month, day := GetDateDetail(date)
	path := CreateFolderDay(day, month, year)

	if format == "md5" {
		pathMd5 := path + NameFile(date, "md5")
		go WriteFile(pathMd5, model.SaveMD5(content), date)
	}
	if format == "sha1" {
		pathSha1 := path + NameFile(date, "sha1")
		go WriteFile(pathSha1, model.SaveSHA1(content), date)
	}
	if format == "sha256" {
		pathSha256 := path + NameFile(date, "sha256")
		go WriteFile(pathSha256, model.SaveSHA256(content), date)
	}
}
