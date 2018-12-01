package main

import (
	"reflect"
	"testing"
)

// func Test1(t *testing.T) {
// 	expected := 15
// 	result, _ := calc("1 2 3 4 + * + =")
// 	if result != expected {
// 		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
// 	}
// }

// func Test2(t *testing.T) {
// 	expected := 21
// 	result, _ := calc("1 2 + 3 4 + * =")
// 	if result != expected {
// 		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
// 	}
// }

// func Test3(t *testing.T) {
// 	expected := -19
// 	result, _ := calc("3 3 / 4 5 * -")
// 	if result != expected {
// 		t.Errorf("results not match\nGot: %v\nExpected: %v", result, expected)
// 	}
// }

// func Test4(t *testing.T) {
// 	expected := "strconv.Atoi: parsing \"g\": invalid syntax"
// 	_, err := calc("g 0 100 500 * / -")
// 	if err.Error() != expected {
// 		t.Errorf("results not match\nGot: %v\nExpected: %v", err.Error(), expected)

// 	}
// }
// func Test5(t *testing.T) {
// 	expected := "strconv.Atoi: parsing \"1.1\": invalid syntax"
// 	_, err := calc("1.1 2.2 3.3 + - /")
// 	if err.Error() != expected {
// 		t.Errorf("results not match\nGot: %v\nExpected: %v", err.Error(), expected)

// 	}
// }

func TestRight(t *testing.T) {
	var tests = []struct {
		expected int
		input    string
	}{
		{15, "1 2 3 4 + * + ="},
		{21, "1 2 + 3 4 + * ="},
		{-19, "3 3 / 4 5 * -"},
	}
	for _, item := range tests {
		result, _ := calc(item.input)
		if !reflect.DeepEqual(result, item.expected) {
			t.Error("expected", item.expected, "got", result)
		}
	}
}

func TestWrong(t *testing.T) {
	var tests = []struct {
		expected string
		input    string
	}{
		{"strconv.Atoi: parsing \"g\": invalid syntax", "g 0 100 500 * / -"},
		{"strconv.Atoi: parsing \"1.1\": invalid syntax", "1.1 2.2 3.3 + - /"},
	}
	for _, item := range tests {
		_, err := calc(item.input)
		if !reflect.DeepEqual(err.Error(), item.expected) {
			t.Error("expected", item.expected, "got", err.Error())
		}
	}
}
