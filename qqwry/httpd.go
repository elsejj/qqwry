package qqwry

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

var db *Db

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func formatFromRequest(r *http.Request) (func(result SearchResult) string, string) {
	format := r.URL.Query().Get("format")
	if format == "" {
		contentType := r.Header.Get("Content-Type")
		if contentType == "application/json" {
			format = "json"
		}
	}
	switch format {
	case "json":
		return FormatJSON, "application/json"
	default:
		return FormatCSV, "text/csv"
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	ips := r.URL.Query()["ip"]
	outputFormatter, contentType := formatFromRequest(r)
	w.Header().Set("Content-Type", contentType)
	flatIps := make([]string, 0, len(ips))
	for _, ip := range ips {
		a := strings.Split(ip, ",")
		flatIps = append(flatIps, a...)
	}
	isArray := contentType == "application/json" && len(flatIps) > 1
	if isArray {
		w.Write([]byte("["))
	}
	isFirst := true
	for _, ip := range flatIps {
		if !isFirst {
			if isArray {
				w.Write([]byte(",\n"))
			} else {
				w.Write([]byte("\n"))
			}
		}
		isFirst = false
		result := SearchIp(db, ip)
		if result.Ip == "" {
			http.Error(w, "IP not found: "+ip, http.StatusNotFound)
			return
		}
		w.Write([]byte(outputFormatter(result)))
	}
	if isArray {
		w.Write([]byte("]"))
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "No body provided", http.StatusBadRequest)
		return
	}
	outputFormatter, contentType := formatFromRequest(r)
	w.Header().Set("Content-Type", contentType)
	SearchReplace(db, r.Body, w, outputFormatter)
}

func normalizeAddress(addr string) (string, []string) {
	a := strings.Split(addr, ":")
	host := "0.0.0.0"
	port := "11223"
	if len(a) == 1 {
		port = a[0]
	} else {
		if len(a[0]) > 0 {
			host = a[0]
		}
		port = a[1]
	}
	addr = host + ":" + port
	if host != "0.0.0.0" {
		return addr, []string{host + ":" + port}
	}
	return addr, enumerateLocalIpAddresses(port)
}

func enumerateLocalIpAddresses(port string) []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	ips := []string{"127.0.0.1:" + port} // Always include localhost
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String()+":"+port)
		}
	}
	return ips
}

func StartHttp(addr, mount, dbPath string) error {
	var err error
	db, err = NewDb(dbPath)
	if err != nil {
		return err
	}
	addr, ips := normalizeAddress(addr)
	for _, a := range ips {
		fmt.Printf("Starting HTTP server on http://%s%s\n", a, mount)
	}
	http.HandleFunc(mount, handler)

	return http.ListenAndServe(addr, nil)
}
