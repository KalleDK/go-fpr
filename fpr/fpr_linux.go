package fpr

import (
	"os"
	"strings"
)

func resolve(d string) string {
	if strings.HasPrefix(d, "$") {
		d = os.Getenv(d[1:len(d)])
	}
	return d
}
