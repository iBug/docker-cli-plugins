package namesgenerator

import (
	"fmt"
	"regexp"
	"strings"
)

// To update the lists, fetch the latest version from:
// https://raw.githubusercontent.com/moby/moby/master/pkg/namesgenerator/names-generator.go

// Export the lists of left and right names
var (
	Left  = left[:]
	Right = right[:]

	ls = strings.Join(Left, "|")
	rs = strings.Join(Right, "|")
	Re = regexp.MustCompile(fmt.Sprintf("^(%s)_(%s)[0-9]?$", ls, rs))
)

func IsRandomName(s string) bool {
	return Re.FindString(s) != ""
}
