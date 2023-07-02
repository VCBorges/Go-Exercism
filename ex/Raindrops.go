package ex

import (
	"strings"
	"fmt"
)

// Raindrops is a function that converts a number to a string, the contents of which depends on the number's factors.

// func Convert(number int) string {
// 	var result string

//     if (number % 3 == 0){
//         result += "Pling"
//     }  
// 	if (number % 5 == 0){
//     	result += "Plang"
//     } 
// 	if (number % 7 == 0){
//     	result += "Plong"
//     } 
// 	if (result == ""){
//     	result = fmt.Sprint(number)
//     }

// 	return result
// }

func Convert(number int) string {
	var sb strings.Builder

	if number%3 == 0 {
		sb.WriteString("Pling")
	}
	if number%5 == 0 {
		sb.WriteString("Plang")
	}
	if number%7 == 0 {
		sb.WriteString("Plong")
	}
	result := sb.String()

	if result == "" {
		result = fmt.Sprint(number)
	}

	return result
}
