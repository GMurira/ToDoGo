package main

import "fmt"

func main() {
	shortGoalng := "Watch Go crash"
	fullGolang := "Watch Hull Course"
	reward := "Reward myself with a Drink"

	taskItems := []string{shortGoalng, fullGolang, reward}

	fmt.Println()
	fmt.Println("Here are the tasks TODO")

	for index, task := range taskItems {
		fmt.Println(index, task)
	}

}
