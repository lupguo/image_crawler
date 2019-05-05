// nodes 提取指定doc文档中的img url
package nodes

import (
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

type nodeVisitFunc func(*html.Node) (links []string)

// ExtractImgUrls 提取页面中的图片URL
func ExtractImgUrls(doc *html.Node) (urls []string) {
	// 去重
	seen := make(map[string]bool)

	// 先序遍历
	links := forEachNode(doc, visitImgNode, nil)
	for _, link := range links {
		if _, ex := seen[link]; !ex && link != "" {
			seen[link] = true
			urls = append(urls, link)
		}
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
