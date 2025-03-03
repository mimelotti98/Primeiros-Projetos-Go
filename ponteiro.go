package main

import "fmt"

func incial(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	incial(&x)
	fmt.Println(x)
}
