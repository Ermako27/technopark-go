package main

import (
	"fmt"
)

// вам необходимо поправить баги в этих функциях чтобы тесты в main_test.go проходили

func ExistCounter() int {
	ids := []int{1, 2, 3, 4, 5, 6}
	enabled := map[int]bool{
		1: false,
		2: true,
		4: false,
		6: true,
	}
	totalExist := 0
	for _, id := range ids {
		if _, i := enabled[id]; i {
			fmt.Println(i)
			totalExist++
		}
	}
	return totalExist
}

func Shadowing() int {
	x := 1
	for i := 0; i <= 10; i++ {
		x = x + 1
		x *= 2
	}
	return x
}

func BadMap() (resultErr error) {
	defer func() {
		if err := recover(); err != nil {
			resultErr = fmt.Errorf("recover: %v", err)
		}
	}()

	users := make(map[int]string)
	users[100500] = "rvasily"
	return nil
}

func main() {
	fmt.Println(ExistCounter()) // 2
	fmt.Println(Shadowing())    // 1
	fmt.Println(BadMap())       // recover: assignment to entry in nil map
}
