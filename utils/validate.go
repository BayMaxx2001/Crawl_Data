package utils

import (
	"regexp"
)

func IsDate(str string) bool {
	re 		:= regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	matched := re.MatchString(str)
	return (matched == true)
}
