package main

import (
	"sort"
	"strconv"
	"strings"
)

// сюда вам надо писать функции, которых не хватает, чтобы проходили тесты в gotchas_test.go

// ReturnInt - 1
func ReturnInt() int {
	num := 1
	return num
}

// ReturnFloat - 2
func ReturnFloat() float32 {
	num := 1.1
	return float32(num)
}

// ReturnIntArray - 3
func ReturnIntArray() [3]int {
	array := [3]int{1, 3, 4}
	return array
}

// ReturnIntSlice - 4
func ReturnIntSlice() []int {
	slice := []int{1, 2, 3}
	return slice
}

// IntSliceToString - 5
func IntSliceToString(slice []int) string {
	strSlice := make([]string, len(slice))

	for i, n := range slice {
		strSlice[i] = strconv.Itoa(n)
	}

	str := strings.Join(strSlice, "")
	return str
}

// MergeSlices - 6
func MergeSlices(slice1 []float32, slice2 []int32) []int {
	result := []int{}
	for _, val := range slice1 {
		result = append(result, int(val))
	}

	for _, val := range slice2 {
		result = append(result, int(val))
	}
	return result
}

// GetMapValuesSortedByKey - 7
func GetMapValuesSortedByKey(items map[int]string) []string {
	result := []string{}
	sortedKeys := []int{}
	for key := range items {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Ints(sortedKeys)

	for _, val := range sortedKeys {
		result = append(result, items[val])
	}

	return result
}
