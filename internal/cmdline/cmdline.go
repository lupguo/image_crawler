package cmdline

import (
	"flag"
)

const (
	AnalyzeByNode  = "node"
	AnalyzeByRegex = "regex"
)

var (
	testData = map[string]string{
		"pageUrl":  "https://blog.golang.org/survey2018-results",
		"storage":  "/tmp",
		"analyzed": AnalyzeByRegex,
	}
	StDir    string
	PgUrl    string
	Analyzed string
)

func init() {
	flag.StringVar(&PgUrl, "url", testData["pageUrl"], "page url request by crawler")
	flag.StringVar(&StDir, "d", testData["storage"], "download image storage dirname")
	flag.StringVar(&Analyzed, "analyzed", testData["analyzed"], "url page analyzed method (node|regex)")
	flag.Parse()
}
