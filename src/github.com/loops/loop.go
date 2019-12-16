package main

import "fmt"

func main() {
	// slice - unordered list
	coursesInProg := []string{
		"Docker deep dive",
		"Docker Clustering",
		"Docker and Kubernetes"}
	coursesCompleted := []string{
		"Docker deep dive",
		"Go fundamentals",
		"Puppet fundamentals",
	}

	for i, course := range coursesInProg {
		fmt.Println(i, course)
		for _, completed := range coursesCompleted {
			if completed == course {
				fmt.Println("hey, we found a clash:\n", completed)
			}
		}
	}
}
