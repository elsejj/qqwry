package qqwry

import "regexp"

var ipV4Regex = regexp.MustCompile(`"?(\d{1,3}\.){3}\d{1,3}"?`)

func IsIpV4(ip string) bool {
	if len(ip) < 7 || len(ip) > 15 {
		return false
	}
	return ipV4Regex.MatchString(ip)
}
