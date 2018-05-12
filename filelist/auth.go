package filelist

import (
	"os"
	"errors"
)

func getFLAuth() (string, string, error) {
	u := os.Getenv("FLBROWSER_USER")
	p := os.Getenv("FLBROWSER_PWD")

	if len([]rune(u)) == 0 || len([]rune(p)) == 0{
		return "", "", errors.New("Missing user/pwd; set FLBROWSER_USER/FLBROWSER_PWD")
	}

	return u, p, nil
}
