package qqwry

import (
	"encoding/json"
	"net/http"
)

var db *Db

type result struct {
	Ip, Country, Area string
}

func handler(w http.ResponseWriter, r *http.Request) {
	var resp result
	iplist, ok := r.URL.Query()["ip"]
	if ok {
		w.Write([]byte("[\n"))
		for _, ip := range iplist {
			resp.Ip = ip
			resp.Country, resp.Area = db.Search(ip)
			buff, _ := json.Marshal(&resp)
			w.Write(buff)
			w.Write([]byte(",\n"))
		}
		w.Write([]byte("]"))
	}
}

func StartHttp(addr, dbpath string) error {
	var err error
	db, err = NewDb(dbpath)
	if err != nil {
		return err
	}
	http.HandleFunc("/", handler)
	return http.ListenAndServe(addr, nil)
}
