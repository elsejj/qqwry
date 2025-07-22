package qqwry

import (
	"testing"
)

var dbpath = "qqwry.dat"
var ip = "183.224.52.133"

func TestSearch(t *testing.T) {
	db, _ := NewDb(dbpath)
	c, a := db.Search(ip)
	t.Log(c, a)
}

func BenchmarkSearch(b *testing.B) {
	db, _ := NewDb(dbpath)
	for i := 0; i < b.N; i++ {
		db.Search(ip)
	}
}

func TestUpdate(t *testing.T) {
	//Update(true)
}

func TestDump(t *testing.T) {
	db, _ := NewDb(dbpath)
	db.Dump()
}

/*
func TestHttpd(t *testing.T) {
	err := StartHttp("127.0.0.1:8000", dbpath)
	if err != nil {
		t.Log(err)
	}
}
*/
