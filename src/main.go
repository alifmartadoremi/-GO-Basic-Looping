package main

import(
	"fmt"
)

func main() {
	var offset int

	fmt.Print("Masukkan jumlah looping : ")
	fmt.Scanln(&offset)

	for i := 0; i<=offset; i++{
		for x := 0; x<=i; x++{
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
}