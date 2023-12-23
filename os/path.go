package xos

import (
	"os"
	"strings"
)

func MakePath(s ...string) string {
	u := []string {}

	for _, t := range s {
		if t != "" {
			u = append(u, t)
		}
	}

	return strings.Join(u, string(os.PathListSeparator))
}
