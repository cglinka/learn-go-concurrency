package main

import (
	"fmt"
)

func main() {
	ar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := 0; i < len(ar); i += 5 {
		fmt.Println(ar[i : i+5])
	}
}
