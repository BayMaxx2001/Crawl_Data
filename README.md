# Task 1: Crawl Data
## Problem
1. Request https://malshare.com/daily  
2. Save all information in this link to your device.

## Getting Started
First, we must installation Goquerry.
```
go get github.com/PuerkitoBio/goquery
```
Second, run file `main.go` crawl data to your device.
```
go run main.go
```
## Steps
1. Request to https://malshare.com/daily.
2. Extract HTML response and get date data by get element table -> tr -> td.
3. Classify data is the md5, sha1,sha256 save to map.
4. Create and write the file.

*Note: Because the data set is too big, `output` example as the first 5 days.*
