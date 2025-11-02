package main

import (
	"fmt"
	"time"
)

func main() {
	// var age int32
	// _, ageErr := fmt.Scan(&age)
	// if ageErr != nil {
	// 	fmt.Printf("can't read age.")
	// 	return
	// }
	// var message string = "hello %s you have %d years old"
	// var name = "amir"
	// var allow = age > 18

	// if allow {
	// 	fmt.Print("you can have beer \n")
	// } else {
	// 	fmt.Println("you can't have beer")
	// }
	// fmt.Printf(message, name, age)
	// var i int32
	// _, iErr := fmt.Scan(&i)
	// if iErr != nil {
	// 	fmt.Printf("can't read age.")
	// 	return
	// }
	// if i == 1 {
	// 	fmt.Println("one")
	// } else if i == 2 {
	// 	fmt.Println("two")
	// }

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// }

	dateNow := time.Now()
	weekday := (dateNow.Weekday().String())
	switch weekday {
	case "Friday":
		fmt.Println("to day is working date")
	case "Saturday ":
		fmt.Println("to day is weekend")
	case "sunday ":
		fmt.Println("to day is weekend")
	}

}
