package main

import (
	"app/ex"
	"fmt"
)



func main()  {
	distance, err := ex.Distance("GGACGGATTCTG", "AGGACGGATTCT")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(distance)
}