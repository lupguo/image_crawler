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
	}
	StDir       string
	PgUrl       string
	Analyzed    string
	Concurrence int
	Sleep		int
)

func init() {
	flag.StringVar(&PgUrl, "url", testData["pageUrl"], "page url request by crawler")
	flag.StringVar(&StDir, "d", testData["storage"], "download image storage dirname")
	flag.StringVar(&Analyzed, "analyzed", AnalyzeByRegex, "url page analyzed method (node|regex)")
	flag.IntVar(&Concurrence, "c", 20, "the concurrence number of image crawler")
	flag.IntVar(&Sleep, "sleep", 0, "sleep time (in ms), reduce the rate of image request on the premise of concurrent request of image crawler")
	flag.Parse()
}
