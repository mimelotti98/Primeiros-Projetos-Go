package main

import "fmt"

func main() {
	//for e If
	for i := 0; i <= 10; i++ {

		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}

		//Switch

		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("Um")
		case 2:
			fmt.Println("Dois")
		case 3:
			fmt.Println("Três")
		case 4:
			fmt.Println("Quatro")
		case 5:
			fmt.Println("Cinco")
		}
	}

}
