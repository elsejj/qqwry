// the qqwry.dat file format is introduced in http://blog.csdn.net/cnss/article/details/77628
// the qqwry.dat file can be download from http://www.cz88.net/

package qqwry

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toOffset3(b []byte, o int) int {
	return int(b[o+0]) | int(b[o+1])<<8 | int(b[o+2])<<16
}

func toOffset4(b []byte, o int) int {
	return int(b[o+0]) | int(b[o+1])<<8 | int(b[o+2])<<16 | int(b[o+3])<<24
}

func ip2str(ip []byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[3], ip[2], ip[1], ip[0])
}

func parseIp(ip string) (int, error) {
	a := strings.Split(ip, ".")
	if len(a) != 4 {
		return 0, errors.New("invalid ip format")
	}
	b := make([]byte, 4)

	for i := 0; i < 4; i++ {
		n, e := strconv.Atoi(a[3-i])
		if e != nil {
			return 0, e
		}
		b[i] = byte(n)
	}
	return toOffset4(b, 0), nil
}

// the qqwry.dat database
type Db struct {
	qqwry []byte
}

func NewDb(path string) (*Db, error) {
	db := new(Db)
	err := db.doInit(path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (self *Db) doInit(path string) error {
	if self.qqwry != nil {
		return nil
	}
	fi, err := os.Lstat(path)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	self.qqwry = make([]byte, fi.Size())

	_, err = file.Read(self.qqwry)
	if err != nil {
		self.qqwry = nil
		return err
	}

	return nil
}

func (self *Db) readIp(o int) (start, end, country, area []byte) {
	qqwry := self.qqwry
	start = qqwry[o : o+4]
	o = toOffset3(qqwry, o+4)
	end = qqwry[o : o+4]

	country, area = self.readInfo(o + 4)
	return start, end, country, area
}

func (self *Db) readText(o int) (int, []byte) {
	qqwry := self.qqwry

	if qqwry[o] == 0x02 {
		o = toOffset3(qqwry, o+1)
		_, t := self.readText(o)
		return 4, t
	} else {
		e := o
		for ; qqwry[e] != 0; e++ {
		}
		return e - o + 1, qqwry[o:e]
	}
}

func (self *Db) readInfo(o int) (country, area []byte) {
	qqwry := self.qqwry

	if qqwry[o] == 0x01 {
		o = toOffset3(qqwry, o+1)
		return self.readInfo(o)
	}
	n, country := self.readText(o)
	_, area = self.readText(o + n)
	return country, area
}

// dump db contents as csv format to console
func (self *Db) Dump() {
	qqwry := self.qqwry

	o := toOffset4(qqwry, 0)
	e := toOffset4(qqwry, 4)

	for ; o < e; o += 7 {
		start, end, country, area := self.readIp(o)
		fmt.Println(ip2str(start), ip2str(end), GbkString(country), GbkString(area))
	}
}

// search a ip , return it's country and area when found
// return "Unknown Unknown" when not found
func (self *Db) Search(ip string) (string, string) {
	qqwry := self.qqwry

	nip, err := parseIp(ip)
	if err != nil {
		return "Unknown", "Unknown"
	}
	o := toOffset4(qqwry, 0)
	e := toOffset4(qqwry, 4)

	last := (e - o) / 7
	first := 0
	mid := (last - first) / 2
	for {
		val := toOffset4(qqwry, o+mid*7)
		if nip > val {
			first = mid
			mid = first + (last-first)/2
		}
		if nip < val {
			last = mid
			mid = first + (last-first)/2
		}
		if nip == val {
			break
		}
		if mid == first {
			break
		}
	}

	_, _, country, area := self.readIp(o + mid*7)
	return GbkString(country), GbkString(area)
}
