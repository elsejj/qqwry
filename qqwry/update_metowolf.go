package qqwry

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// see https://github.com/metowolf/qqwry.dat
var metowolfUrl = "https://github.com/metowolf/qqwry.dat/releases/latest/download/qqwry.dat"

func UpdateMetowolfQQWry(saveTo string) error {

	httpClient := &http.Client{}

	resp, err := httpClient.Get(metowolfUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(saveTo)
	if err != nil {
		return err
	}
	defer file.Close()

	totalSize := resp.ContentLength
	fmt.Printf("Total size: %d bytes\n", totalSize)
	chunkSize := int64(100 * 1024) // 100 KB
	var downloaded int64 = 0
	buffer := make([]byte, chunkSize)
	for {
		n, err := resp.Body.Read(buffer)
		if n > 0 {
			downloaded += int64(n)
			file.Write(buffer[:n])
			fmt.Printf("\rDownloaded %d bytes (%.2f%%)", downloaded, float64(downloaded)/float64(totalSize)*100)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	fmt.Println("")
	return nil
}
