package parser

import "strings"

func CleanEscSequence(data string) string {
	if len(data) > 0 {
		return strings.Replace(data, "\r\n", "", -1)
	}
	return ""
}
