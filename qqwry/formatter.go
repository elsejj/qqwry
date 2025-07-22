package qqwry

import "fmt"

func FormatCSV(r SearchResult) string {
	return fmt.Sprintf("%s %s %s %s %s", r.Ip, r.Country, r.Province, r.City, r.Operator)
}

func FormatJSON(r SearchResult) string {
	return fmt.Sprintf(`{"ip":"%s","country":"%s","province":"%s","city":"%s","operator":"%s"}`,
		r.Ip, r.Country, r.Province, r.City, r.Operator)
}
