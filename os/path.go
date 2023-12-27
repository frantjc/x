package xos

import (
	"os"
	"strings"

	xslice "github.com/frantjc/x/slice"
)

func JoinPath(s ...string) string {
	return strings.Join(
		xslice.Filter(s, func(t string, i int) bool {
			return t != "" && i == xslice.IndexOf(s, t)
		}),
		string(os.PathListSeparator),
	)
}
