package main
import "fmt"

func main() {
	name := []string{"Nikhil","Bhanushali","Don","Someone"}
	fmt.Println(name)
	name = append(name[:2],name[3:]...)
	fmt.Println(name)
}
