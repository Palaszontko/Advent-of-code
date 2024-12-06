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

func SliceSum2D(slice [][]int) int {
	amount := 0

	for _, row := range slice {
		amount += SliceSum(row)
	}

	return amount
}

func Copy2DSlice(original [][]byte) [][]byte {
	copySlice := make([][]byte, len(original))
	for i := range original {
		copySlice[i] = make([]byte, len(original[i]))
		copy(copySlice[i], original[i])
	}
	return copySlice
}
