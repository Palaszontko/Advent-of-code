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

func DeepCopy2D[T any](matrix [][]T) [][]T {
	if matrix == nil {
		return nil
	}

	result := make([][]T, len(matrix))
	for i := range matrix {
		if matrix[i] != nil {
			result[i] = make([]T, len(matrix[i]))
			copy(result[i], matrix[i])
		}
	}
	return result
}

func CountIn2DSlice[T any](slice [][]T, condition func(T) bool) int {
	amount := 0
	for _, row := range slice {
		for _, val := range row {
			if condition(val) {
				amount += 1
			}
		}
	}

	return amount
}

func Contains2DSlice(slice [][]int, condition func(int) bool) bool {
	for _, row := range slice {
		for _, val := range row {
			if condition(val) {
				return true
			}
		}
	}

	return false
}

func InsertAtIndexIn2DSlice[S [][]E, R []E, E any](slice2d S, row R, index int) S {
	return append(slice2d[:index], append([][]E{row}, slice2d[index:]...)...)
}

func Gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int64) int64 {
	return a * b / Gcd(a, b)
}

func LcmMultipleNumbers(numbers []int64) int64 {
	result := numbers[0]
	for _, num := range numbers[1:] {
		result = Lcm(result, num)
	}
	return result
}
