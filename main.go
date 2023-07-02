package main

import (
	"fmt"
	"app/ex"
)



func main()  {
	con, _ := ex.CollatzConjecture(12)
	fmt.Println(con)
}