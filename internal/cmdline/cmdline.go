package cmdline

import (
	"flag"
)

var (
	testData = map[string]string{
		"pageUrl": "https://blog.golang.org/survey2018-results",
		"storage": "/tmp",
	}
	StDir string
	PgUrl string
)

func init() {
	flag.StringVar(&PgUrl, "url", testData["pageUrl"], "page url request by crawler")
	flag.StringVar(&StDir, "d", testData["storage"], "download image storage dirname")
	flag.Parse()
}
