// httpImg download http image file
package httpImg

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Download download url file into stDir
func Download(realurl, dir string) (dstfile string, err error) {
	resp, err := http.Get(realurl)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// check whether Content-Type is prefix image/
	ctype := resp.Header.Get("Content-Type")
	if ! strings.HasPrefix(ctype, "image/") {
		return "", errors.New(realurl + " is not an image file")
	}

	// confirm new filename
	basename := filepath.Base(realurl);
	if index := strings.IndexAny(basename, "?#"); index != -1 {
		basename = basename[:index]
	}

	dstfile = fmt.Sprintf("%s/%s", path.Clean(dir), basename)
	fi, err := os.Create(dstfile)
	if err != nil {
		return
	}
	defer fi.Close()

	// copy image file
	_, err = io.Copy(fi, resp.Body)
	if err != nil {
		return
	}

	return
}
