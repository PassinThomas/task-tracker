package util

import (
	"fmt"
	"strings"
	"errors"
)

var (Verbose bool)

func Vlog(v bool, s string) {
	if v == true {
		fmt.Println(s)
	}
	return
}

func ParseStr(s string) error {
	if strings.TrimSpace(s) == "" {
		return errors.New("Input cannot be empty.")
	}
	if len([]rune(s)) > 255 {
        return errors.New("Title too long")
    }
	if strings.ContainsAny(s, "/\\") {
        return errors.New("Title cannot contain slashes")
    }
	return nil
}