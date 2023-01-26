package main
import "fmt"

func controls() {

	fmt.Print("enter a number = ")
	var input int
	fmt.Scanln(&input)

	if(input % 2 == 0) {
		fmt.Println("hey you are even")
	} else {
		fmt.Println("hey you are odd")
	}

	if x := 10 ; x%2 == 0 {
		fmt.Println("hey you are even")
	} else {
		fmt.Println("hey you are odd")
	}

	data := 10
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
	
	case 100:
		data = 10007
		fmt.Println(data)
		fallthrough

	case 11:
		data = 104
		fmt.Println(data)
		fallthrough

	case 15:
		data = 1002
		fmt.Println(data)
	}

}