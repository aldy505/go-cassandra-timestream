package main

import "strings"

func Quote(s string) string {
	return "'" + strings.ReplaceAll(s, "'", "\\'") + "'"
}
