package main

import (
	"errors"
	"log"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattn/go-isatty"
)

func FatalCheck(err error) {
	if err != nil {
		log.Fatalf("%v", err)
		panic(err)
	}
}

func FilterIPV4(ips []net.IP) []string {
	var ret = make([]string, 0)
	for _, ip := range ips {
		if ip.To4() != nil {
			ret = append(ret, ip.String())
		}
	}
	return ret
}

func MkdirIfNotExist(folder string) error {
	if _, err := os.Stat(folder); err != nil {
		if err = os.MkdirAll(folder, 0700); err != nil {
			return err
		}
	}
	return nil
}

func DirExists(folder string) bool {
	_, err := os.Stat(folder)
	return err == nil
}

func DisplayProgressBar() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) && displayProgress
}

func FolderOf(url string) string {
	safePath := filepath.Join(os.Getenv("HOME"), dataFolder)
	fullyQualifiedPath, err := filepath.Abs(filepath.Join(safePath, filepath.Base(url)))
	FatalCheck(err)

	// must ensure fully qualified path is child of safe path
	// to prevent directory traversal attack
	// using Rel function to get relative between parent and child
	// if relative join base == child, then child path MUST BE real child
	relative, err := filepath.Rel(safePath, fullyQualifiedPath)
	FatalCheck(err)

	if strings.Contains(relative, "..") {
		FatalCheck(errors.New("you may be a victim of directory traversal path attack"))
		return "" // return is redundant because in fatal check we have panic, but compiler is not able to check
	} else {
		return fullyQualifiedPath
	}
}

func TaskFromUrl(url string) string {
	// task is just download file name
	// so we get download file name on url
	filename := filepath.Base(url)
	return filename
}

func IsUrl(s string) bool {
	_, err := url.Parse(s)
	return err == nil
}
