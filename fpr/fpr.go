package fpr

import (
	fp "path/filepath"
	"strings"
)

func Resolve(d string) string {
	sep := string(fp.Separator)
	parts := []string{}
	for _, part := range strings.Split(fp.FromSlash(d), sep) {
		parts = append(parts, resolve(part))
	}
	return fp.Clean(strings.Join(parts, sep))
}
