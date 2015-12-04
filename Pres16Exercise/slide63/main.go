package main
import "fmt"

func printFriends(names ...string) {
	for _,val := range names{
		fmt.Println(val)
	}
}

func main() {
	friends := []string{"Viren","Sewag","Sachin","James","Tom"}
	printFriends(friends...)
}
