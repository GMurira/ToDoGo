package main

import "fmt"

func main() {
	fmt.Println()
	var fullGolang = "Watch Hull Course"
	var reward = "Reward myself with a Drink"
	var taskItems = []string{fullGolang, reward}

	printTasks(taskItems)
	fmt.Println()
	taskItems = addTasks(taskItems, "Go for a run")
	fmt.Println("Updated List")
	printTasks(taskItems)
}

func printTasks(taskItems []string) {
	fmt.Println("Here are the tasks TODO")

	for index, task := range taskItems {
		//fmt.Println(index+1, ":", task)
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTasks(taskItems []string, newTask string) []string {
	var updatedTasks = append(taskItems, newTask)
	return updatedTasks
}
