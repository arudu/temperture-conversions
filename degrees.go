package main

import "fmt"


func degrees()  {

	fmt.Print("Enter a temp in fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := ((input - 32) * 5/9)
	fmt.Println("The temp in celcius is:", output, "°C")
}

func main() {
	degrees()
}
