package utils

import (
	"fmt"
	"os"
	"strconv"
)

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("failed to convert %q to int: %v", s, err))
	}
	return i
}

func ReadFile(path string) string {
	lines, err := os.ReadFile(path)

	if err != nil {
		panic(fmt.Sprintf("failed to read file %q: %v", path, err))
	}

	return string(lines[:len(lines)-1])
}

func StringToIntSlice(ss []string) []int {
	intValues := make([]int, len(ss))
	for i, s := range ss {
		intValues[i] = MustAtoi(s)
	}
	return intValues
}

func SliceSum(slice []int) int {
	amount := 0

	for _, val := range slice {
		amount += val
	}

	return amount
}
