package db

import "strings"

func ToTSQuery(src string) string {
	if len(src) == 0 {
		return ""
	}
	return strings.Join(strings.Fields(src), ":* & ") + ":*"
}
