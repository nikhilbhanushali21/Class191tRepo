package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		for i%2 == 0 {
			for{
				fmt.Println(i)
				break
			}
		}
	}
}
