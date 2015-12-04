package main
import "fmt"

func main() {
	ages := map[string]int{
		"Nikhil":26,
		"Bhanushali":27,
		"Tushar":24,
		"Suarabh":1,
		"Ajay":23,
	}
	ages["Ajanth"]=14
	ages["Broil"]=26
	if _, exists := ages["Nikhil"]; exists {
		delete(ages, "Nikhil")
	}
	for key,val := range ages{
		fmt.Println(key,"is",val,"years old.")
	}
	fmt.Println("Length of map is",len(ages))
}
