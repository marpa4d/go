// Для понимания, как достигается сложность алгоритмов O(n) и операций над данными. Примеры.
package main

import (
	"fmt"
	"time"
) //TIP S

func On(n int) {

	timer := time.Now()

	count := 0
	for i := 0; i < n; i++ {
		count++
	}
	fmt.Println("O(n):", time.Since(timer))
}

func main() {

	n := 10000000 // O(1)

	// On - сложность O(n)
	On(n)
}
