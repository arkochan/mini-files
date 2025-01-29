package main

import "strings"

func GetBaseFileName(filepath string) string {
	lastInd := strings.LastIndex(filepath, "_")
	return filepath[:lastInd]
}
