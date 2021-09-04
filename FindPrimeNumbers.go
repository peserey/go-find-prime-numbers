package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scanf("%d", &number)
	for i := 2; i <= number; i++ {
		var myFlag bool
		myFlag = false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				myFlag = true
			}
		}
		if myFlag == false {
			fmt.Println(i)
		}
	}
}
