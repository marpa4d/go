// Для понимания, как достигается сложность алгоритмов O(n) и операций над данными. Примеры.
package main

import (
	"fmt"
	"time"
)

// On - сложность O(n)!
func On(n int) {

	timer := time.Now()

	count := 0
	for i := 0; i < n; i++ {
		count++
	}
	fmt.Println("O(n):", time.Since(timer))
}

// On2 - сложность O(n^2)
func On2(n int) {

	timer := time.Now()

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count++
		}
	}
	fmt.Println("O(n^2):", time.Since(timer))

}

func main() {

	n := 100000 // O(1)

	// On - сложность O(n)
	On(n)

	// On2 - сложность O(n^2)
	On2(n)
}
