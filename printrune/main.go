package main

import (
	"fmt"
	"os"
)

func main(){
	for _,arg1 := range os.Args[1:] {
		for _,r := range arg1 {
			fmt.Printf("\\u%04X",r)
		}
		fmt.Println()
	}
}
