package main

import (
	"os"
	"strconv"
	"strings"
)

func GetNumberFromFile(file string) (int, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return -1, err
	}

	return strconv.Atoi(strings.ReplaceAll(string(content), "\n", ""))
}