package main

import (
	"errors"
)

func minLength(text string, minLen int) error {
	length := len(text)
	if length < minLen {
		return errors.New("Text length is shoter that expected")
	}
	return nil
}
