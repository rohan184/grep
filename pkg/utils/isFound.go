package utils

import (
	"strings"
)

func isFound(txt, a string) (string, bool) {
	if strings.Contains(txt, a) {
		return txt, true
	} else {
		return "", false
	}
}
