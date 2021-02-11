package ipfs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gnames/gnsys"
)

type GetterFile struct{}

func (gf GetterFile) Get(source, dest string) error {
	var src, dst *os.File
	exists, err := gnsys.FileExists(source)
	if !exists {
		return fmt.Errorf("file '%s' does not exist", source)
	}

	if err == nil {
		src, err = os.Open(source)
	}

	if err == nil {
		defer src.Close()
		dst, err = os.Create(dest)
	}

	if err == nil {
		defer dst.Close()
		_, err = io.Copy(dst, src)
	}
	return err
}

type GetterHTTP struct{}

func (gh GetterHTTP) Get(source, dest string) (err error) {
	resp, err := http.Get(source)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dst, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, resp.Body)
	return err
}

func isFile(s string) bool {
	res, err := gnsys.FileExists(s)
	if err != nil {
		return false
	}
	return res
}

func isURL(s string) bool {
	res := strings.HasPrefix(s, "https://") || strings.HasPrefix(s, "http://")
	return res
}

func isPathIPNS(s string) bool {
	res := strings.HasPrefix(s, "/ipns/")
	return res
}

func isKeyIPNS(s string) bool {
	res := len(s) > 40 && strings.HasPrefix(s, "k5") &&
		!strings.Contains(s, "/\\")
	return res
}

func isPathIPFS(s string) bool {
	res := len(s) > 40 && strings.HasPrefix(s, "/ipfs/")
	return res
}

func isCID(s string) bool {
	res := len(s) > 40 &&
		strings.HasPrefix(s, "Qm") || strings.HasPrefix(s, "baf") &&
		!strings.Contains(s, "/\\")
	return res
}
