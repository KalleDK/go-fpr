package fpr

import (
	"os"
	"strings"
)

func resolve(d string) string {
	if strings.HasPrefix(d, "%") && strings.HasSuffix(d, "%") {
		d = os.Getenv(d[1 : len(d)-1])
	}
	return d
}
