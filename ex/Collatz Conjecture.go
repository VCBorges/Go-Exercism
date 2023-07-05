package ex

import (
	"fmt"
)

// type teste struct{
// 	name string
// }
// dssfds
func CollatzConjecture(n int) (int, error) {

	if n <= 0 {
		return 0, fmt.Errorf("invalid input")
	}

	var count int
	for n != 1{
		if n % 2 == 0 {
			n /= 2
		} else{
			n = (n * 3) + 1
		}
		count ++
	}

	return count, nil
}

