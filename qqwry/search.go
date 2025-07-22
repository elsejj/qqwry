package qqwry

import (
	"bufio"
	"io"
	"log"
	"strings"
)

type SearchResult struct {
	Ip       string
	Country  string
	Province string
	City     string
	Operator string
}

func searchIp(db *Db, ip string) SearchResult {

	area, operator := db.Search(ip)

	areDetails := strings.Split(area, "–")

	result := SearchResult{
		Ip:       ip,
		Operator: operator,
	}

	if len(areDetails) > 0 {
		result.Country = areDetails[0]
	}
	if len(areDetails) > 1 {
		result.Province = areDetails[1]
	}
	if len(areDetails) > 2 {
		result.City = areDetails[2]
	}
	return result
}

func SearchIp(db *Db, ip string) SearchResult {
	return searchIp(db, ip)
}

func SearchReaders(db *Db, rd io.Reader) []SearchResult {
	var results []SearchResult
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		ip := strings.TrimSpace(scanner.Text())
		if IsIpV4(ip) {
			results = append(results, searchIp(db, ip))
		} else {
			log.Println("Invalid IP format:", ip)
		}
	}
	return results
}

func SearchReplace(db *Db, r io.Reader, w io.Writer, formatter func(SearchResult) string) {
	scanner := bufio.NewScanner(r)
	writer := bufio.NewWriter(w)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Bytes()
		line = ipV4Regex.ReplaceAllFunc(line, func(ip []byte) []byte {
			ipStr := string(ip)
			start := 0
			end := len(ipStr)
			if strings.HasPrefix(ipStr, "\"") {
				start = 1
			}
			if strings.HasSuffix(ipStr, "\"") {
				end = end - 1
			}
			result := searchIp(db, ipStr[start:end])
			return []byte(formatter(result))
		})
		writer.Write(line)
		writer.WriteByte('\n')
	}
}
