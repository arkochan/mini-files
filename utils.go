package main

import "strings"

func GetBaseFileName(filepath string) string {
	lastInd := strings.LastIndex(filepath, ".")
	return filepath[:lastInd]
}
