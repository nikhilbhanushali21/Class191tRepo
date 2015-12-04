package main
import "fmt"

func run(f func(string), names ...string) {
	for _, val := range names{
		f(val)
	}
}

func main() {
	names := []string{"Nikhil","Bhanushali","Thomas","Jerry"}
	greet := func(name string){fmt.Println("Hello,",name)}
	run(greet,names...)
}
