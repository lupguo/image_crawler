// nodes 提取指定doc文档中的img url
package nodes

import (
	"github.com/tkstorm/image_crawler/internal/cmdline"
	"github.com/tkstorm/image_crawler/internal/helper"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type nodeVisitFunc func(*html.Node) (links []string)

// ExtractImgUrls 基于指定类型提取页面中的图片URL
func ExtractImgUrls(resp *http.Response, analyzed string) (urls []string) {
	switch analyzed {
	case cmdline.AnalyzeByNode:
		urls = analyzeByNode(resp)
	case cmdline.AnalyzeByRegex:
		urls = analyzeByRegex(resp)
	}

	return
}

// byHtmlNode 通过分析Html节点获取url信息
func analyzeByNode(resp *http.Response) (urls []string) {
	doc, err := html.Parse(resp.Body)
	if err != nil {
		helper.ErrorOut(err, "parse html doc")
	}

	// 去重 && 先序遍历
	seen := make(map[string]bool)
	links := forEachNode(doc, visitImgNode, nil)
	for _, link := range links {
		if _, ex := seen[link]; !ex && link != "" {
			seen[link] = true
			urls = append(urls, link)
		}
	}
	return
}

// analyzeByRegex 通过正则检测
func analyzeByRegex(resp *http.Response) (urls []string) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		helper.ErrorOut(err, "analyzeByRegex ioutil read error")
	}

	re := regexp.MustCompile(`(?i)(/[^"']*?\.(jpg|jpeg|png|gif|webp|svg))`)
	for _, match := range re.FindAll(body, -1) {
		urls = append(urls, string(match))
	}

	return
}

// 遍历DOM树节点
func forEachNode(node *html.Node, preVisit, postVisit nodeVisitFunc) (links []string) {
	if preVisit != nil {
		if preLinks := preVisit(node); preLinks != nil {
			links = append(links, preLinks...)
		}
	}
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		nxtlinks := forEachNode(n, preVisit, postVisit)
		if len(nxtlinks) == 0 {
			continue
		}
		links = append(links, nxtlinks...)
	}
	if postVisit != nil {
		if postLinks := postVisit(node); postLinks != nil {
			links = append(links, postLinks...)
		}
	}
	return
}

// visitImgNode parse a html node, if is url node, return src attribute,
// else return error
func visitImgNode(node *html.Node) (links []string) {

	// filter not <img /> node
	if node.Type != html.ElementNode || node.Data != "img" {
		return
	}

	// parse <img /> node
	// exts := []string{".jpg", ".png", ".gif", ".jpeg", ".webp", ".svg", ".bmp"}
	for _, attr := range node.Attr {
		if attr.Key == "src" || strings.Contains(attr.Key, "data") {
			imgURL, err := url.Parse(attr.Val)
			if err == nil {
				links = append(links, imgURL.String())
			}
			//if err == nil && inArray(path.Ext(imgURL.Path), exts) {
			//	links = append(links, imgURL.String())
			//}
		}
	}

	return links
}

// inArray check the substr whether in the strArr slice
func inArray(substr string, strArr []string) bool {
	for i := range strArr {
		if strArr[i] == substr {
			return true
		}
	}

	return false
}
