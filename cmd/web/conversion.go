package main

import (
	"regexp"
	"strings"
)

func createSlug(title string) string{
	lowerTitle := strings.ToLower(title)
	reg := regexp.MustCompile(`[^a-z0-9\s]`)
	lowerTitle = string(reg.ReplaceAll([]byte(lowerTitle), []byte("")))
	reg = regexp.MustCompile("[ ]")
	slug := string(reg.ReplaceAll([]byte(lowerTitle), []byte("-")))
	return strings.Trim(slug,"-")
}


