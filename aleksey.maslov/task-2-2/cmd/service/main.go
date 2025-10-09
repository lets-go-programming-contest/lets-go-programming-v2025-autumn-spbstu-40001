package main

import "fmt"

func main() {
	var dishesCount int

	_, err := fmt.Scanln(&dishesCount)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}
	//create heap
	for range dishesCount {
		var rating int

		_, err = fmt.Scanln(&rating)
		if err != nil {
			fmt.Println("Invalid input", err)

			return
		}
		//push
	}

	var selectedDish int

	_, err = fmt.Scanln(&selectedDish)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}

	result := //func
		fmt.println(result)
}
