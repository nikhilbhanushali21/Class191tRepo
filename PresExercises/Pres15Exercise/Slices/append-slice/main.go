package main
import "fmt"

func main() {
	name1 := []string{"Nikhil","Bhanushali"}
	name2 := []string{"Someone","Don"}
	name := append(name1,name2...)
	fmt.Println(name)
}
