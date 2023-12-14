package main

import (
	"errors"
	"strconv"
)

var (
	NegNum = errors.New("negative numbers are not allowed")
)

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", NegNum
	}

	return strconv.FormatInt(int64(num), 2), nil
}
