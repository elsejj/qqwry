package qqwry

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	remoteKeyURL  = "http://update.cz88.net/ip/copywrite.rar"
	remoteDataURL = "http://update.cz88.net/ip/qqwry.rar"
	localKeyURL   = "copywrite.rar"
	localDataURL  = "qqwry.rar"
	dbFileName    = "qqwry.dat"
	tpFileName    = "qqwry.tmp"
)

func urlopen(url string) (*bytes.Buffer, error) {
	var buff bytes.Buffer
	if strings.HasPrefix(url, "http://") {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		io.Copy(&buff, resp.Body)
	} else {
		file, err := os.Open(url)
		if err != nil {
			return nil, err
		}
		io.Copy(&buff, file)
		file.Close()
	}
	return &buff, nil
}

func decode(keyurl, dataurl string) error {
	rk, err := urlopen(keyurl)
	if err != nil {
		return err
	}
	rd, err := urlopen(dataurl)
	if err != nil {
		return err
	}
	var key int32
	for i := 0; i < 6; i++ {
		err = binary.Read(rk, binary.LittleEndian, &key)
		if err != nil {
			return nil
		}
	}

	qqwry := rd.Bytes()
	for i := 0; i < 0x200; i++ {
		key *= 0x805
		key += 1
		key = key & 0xFF
		qqwry[i] = qqwry[i] ^ byte(key)
	}

	file, err := os.Create(tpFileName)
	if err != nil {
		log.Println(err)
		return err
	}

	rd = bytes.NewBuffer(qqwry)
	zr, err := zlib.NewReader(rd)
	if err != nil {
		log.Println(err)
		return err
	}
	io.Copy(file, zr)
	zr.Close()
	file.Close()
	return nil
}

func Update(remote bool) {
	var err error
	if remote {
		err = decode(remoteKeyURL, remoteDataURL)
	} else {
		err = decode(localKeyURL, localDataURL)
	}
	if err == nil {
	}
}
