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
	"golang.org/x/net/html"
	"net/http"
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

	// parse html dom and extract download links
	doc, err := html.Parse(resp.Body)
	if err != nil {
		helper.ErrorOut(err, "parse html doc")
	}
	links := nodes.ExtractImgUrls(doc)

	// goroutine download
	fmt.Printf("total %d images need to be download...\n", len(links))
	ch := make(chan string, len(links))
	for _, dwlink := range links {
		dwurl, err := helper.CorrectUrl(pgUrl, dwlink)
		if err != nil {
			ch <- fmt.Sprintf("%s %s", dwlink, err)
		}

		go func(dwurl string) {
			rs, err := httpImg.Download(dwurl, stDir)
			if err != nil {
				ch <- fmt.Sprintf("%s %s", rs, err)
				return
			}
			ch <- fmt.Sprintf("ok %s => %s", dwurl, rs)
		}(dwurl)
	}

	// output result
	for i := 0; i < len(links); i++ {
		fmt.Println(<-ch)
	}
}
