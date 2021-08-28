package model

type PageInformation struct {
	MD5    string
	SHA1   string
	SHA256 string
	DAY    string
	MONTH  string
	YEAR   string
}

func NewPageInformation(MD5 string, SHA1 string, SHA256 string, DAY string, MONTH string, YEAR string) PageInformation {
	return PageInformation{MD5, SHA1, SHA256, DAY, MONTH, YEAR}
}

func SaveMD5(infor []PageInformation) string {
	var ret string
	for i := range infor {
		ret = ret + infor[i].MD5 + "\n"
	}
	return ret
}

func SaveSHA1(infor []PageInformation) string {
	var ret string
	for i := range infor {
		ret = ret + infor[i].SHA1 + "\n"
	}
	return ret
}

func SaveSHA256(infor []PageInformation) string {
	var ret string
	for i := range infor {
		ret = ret + infor[i].SHA256 + "\n"
	}
	return ret
}
