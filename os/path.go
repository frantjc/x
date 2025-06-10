package xos

import (
	"os"
	"slices"
	"strings"

	xslices "github.com/frantjc/x/slices"
)

func JoinPath(s ...string) string {
	i := -1

	return strings.Join(
		slices.DeleteFunc(s, func(t string) bool {
			i++
			return t == "" || i != xslices.IndexOf(s, t)
		}),
		string(os.PathListSeparator),
	)
}
