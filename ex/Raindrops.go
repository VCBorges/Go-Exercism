package ex

import (
	"strings"
	"fmt"
)


func Convert(number int) string {
	var sb strings.Builder

	if number % 3 == 0 {
		sb.WriteString("Pling")
	}

	if number % 5 == 0 {
		sb.WriteString("Plang")
	}

	if number % 7 == 0 {
		sb.WriteString("Plong")
	}
	
	result := sb.String()

	if result == "" {
		result = fmt.Sprint(number)
	}

	return result
}
