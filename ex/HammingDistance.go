package ex

import (
	"fmt"
)


func Distance(a, b string) (int, error) {
	if len(a) != len(b){
		return 0, fmt.Errorf("different size strings")
	}
	distance := 0
	for i := range a {
		if a[i] != b[i]{
			distance ++
		}
	}
	return distance, nil
}
