package utils

import (
	"strings"
)

// https://malshare.com/daily/yyyy-MM-dd/malshare_fileList.yyyy-MM-dd.all.txt)
func GetAddressOfDate(day string) string {
	address := "https://malshare.com/daily/"
	address  = address + day + "/malshare_fileList." + string(day) + ".all.txt"

	return address
}

// Get day, month, year
func GetDateDetail(day string) (string, string, string) {
	words := strings.Split(day, "-")
	return words[0], words[1], words[2]
}

func GetDateToString(day string, month string, year string) string {
	return year + "-" + month + "-" + day
}
