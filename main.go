// image_crawler 爬取并下载指定URL页面内容中，<img>节点内的图片信息到指定的目录
//
// usage: ./image_crawler -url http://example.com -d /tmp
package main

import (
	"fmt"
	"github.com/tkstorm/image_crawler/internal/cmdline"
	"github.com/tkstorm/image_crawler/internal/helper"
	"github.com/tkstorm/image_crawler/internal/httpImg"
	"github.com/tkstorm/image_crawler/internal/nodes"
	"net/http"
	"sync"
	"time"
)

func main() {
	// input parameter detection
	pgUrl, stDir := cmdline.PgUrl, cmdline.StDir
	if err := helper.MakePath(stDir); err != nil {
		helper.ErrorOut(err, "make path fail")
	}

	// the param of url for crawler using
	resp, err := http.Get(pgUrl)
	if err != nil {
		helper.ErrorOut(err, "get page url")
	}
	defer resp.Body.Close()

	// parse html dom by node and extract download links
	links := nodes.ExtractImgUrls(resp, cmdline.Analyzed)

	// goroutine download
	fmt.Printf("total %d images need to be download...\n", len(links))
	ch := make(chan string)
	sema := make(chan struct{}, cmdline.Concurrence)
	slptime := time.Duration(cmdline.Sleep)
	var wg sync.WaitGroup

	// closer
	go func() {
		wg.Wait()
		close(ch)
	}()

	// worker
	for _, dwlink := range links {
		dwurl, err := helper.CorrectUrl(pgUrl, dwlink)
		if err != nil {
			ch <- fmt.Sprintf("%s %s", dwlink, err)
		}

		wg.Add(1)
		go func(dwurl string) {
			// concurrence setting
			sema <- struct{}{}
			defer func() {
				time.Sleep(slptime * time.Millisecond)
				wg.Done()
				<-sema
			}()

			// download image work
			rs, err := httpImg.Download(dwurl, stDir)
			if err != nil {
				ch <- fmt.Sprintf("%s %s", rs, err)
				return
			}
			ch <- fmt.Sprintf("ok %s => %s", dwurl, rs)
		}(dwurl)
	}

	// output result
	for result := range ch {
		fmt.Println(result)
	}
}
