package main

import("fmt")

func main() {

	var num1, num2 int
	var answer int

	fmt.Print("what numbers would you like to use?")

	fmt.Scan(&num1)
	fmt.Scan(&num2)

	fmt.Print("What operation would you like to use")
	var input string
	fmt.Scan(&input)

	if (input == "add") {

		answer = num1 + num2

	}else if(input == "subtract"){
		answer = num1 - num2

	}else if(input == "multiply"){
		answer = num1 * num2

	}else if(input == "divide"){
		answer = num1 / num2
	}



	fmt.Print(answer)

	



	
}
	
