package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func multiDownload(url, fileName string, start, end, i int) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
	resp, err := http.DefaultClient.Do(req)
	fmt.Println("resp content len", resp.ContentLength, start, end)
	defer resp.Body.Close()
	buf := make([]byte, 30*1024)

	flags := os.O_CREATE | os.O_WRONLY
	realName := fmt.Sprintf("%s-%d", fileName, i)
	fmt.Println("realName:", realName)
	f, err := os.OpenFile(realName, flags, 0777)
	if err != nil {
		log.Fatal("OpenFile fail:", err)
	}

	d, err := io.CopyBuffer(f, resp.Body, buf)
	if err != nil {
		log.Fatal("CopyBuffer fail:", err)
	}
	defer f.Close()
	fmt.Println("d:", d)

}

func costTime(st time.Time) {
	fmt.Println("cost time:", time.Since(st))
}

func download(url string) {
	startTime := time.Now()
	defer costTime(startTime)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	totalSize := int(resp.ContentLength)
	fmt.Println("totalSize:", totalSize)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	filePath := "15-multi-download/test"
	fileNames := strings.Split(url, "/")
	fileName := fmt.Sprintf("%s/%s", filePath, fileNames[len(fileNames)-1])
	os.Mkdir(filePath, 0777)
	start, end, times := 0, 0, 10
	step := totalSize / times
	wg := sync.WaitGroup{}

	for i := 0; i < times; i++ {
		if i == times-1 {
			end = totalSize
		} else {
			end = start + step
		}
		fmt.Println("fileName:", fileName, i, start, end)
		go func(wg *sync.WaitGroup, start, end,i int) {
			defer wg.Done()

			multiDownload(url, fileName, start, end, i)
		}(&wg, start, end,i)
		start = end + 1
		wg.Add(1)
	}
	wg.Wait()
	merge(fileName, times)
}

func merge(fileName string, times int) {
	realFileName, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer realFileName.Close()

	for i := 0; i < times; i++ {
		partFileName := fmt.Sprintf("%s-%d", fileName, i)
		partFile, err := os.Open(partFileName)
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(realFileName, partFile)
		partFile.Close()
		//os.Remove(partFileName)
	}
}

func main() {
	url := "https://down.sandai.net/mac/thunder_4.2.1.65254.dmg"
	download(url)

}
