package helper

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

// MakePath make path if path not exist
func MakePath(dirname string) error {
	if _, err := os.Stat(dirname); err != nil {
		return os.MkdirAll(dirname, 0755)
	}
	return nil
}

// CorrectUrl check the download link, compare with the pageurl, get the correct real url
func CorrectUrl(pageurl string, dwlink string) (realurl string, err error) {
	pgURL, _ := url.Parse(pageurl)

	// begin with real url http(s)://example.com/...
	if strings.HasPrefix(dwlink, "http") {
		return dwlink, nil
	}

	// begin with //example.com/...
	if strings.HasPrefix(dwlink, "//") {
		realurl = fmt.Sprintf("%s:%s", pgURL.Scheme, dwlink)
		return
	}

	// begin with relate path: /...
	return fmt.Sprintf("%s://%s/%s", pgURL.Scheme, pgURL.Host, strings.TrimLeft(dwlink, "/")), nil

}

// ErrorOut put error out
func ErrorOut(err error, mark string) {
	if err != nil {
		fmt.Printf("[%s] %s", mark, err)
		os.Exit(-1)
	}
}
