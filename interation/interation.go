package interation

import "strings"

func Repeat(s string, count int) string {
	var result strings.Builder
	for i := 0; i < count; i++ {
		result.WriteString(s)
	}
	return result.String()
}
