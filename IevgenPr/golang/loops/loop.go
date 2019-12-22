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

	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			//continue
			break
		}
	}

	for timer := 3; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Ka boom!")
			//continue
			break
		}
	}

	for i, course := range coursesInProg {
		fmt.Println(i, course)
		for _, completed := range coursesCompleted {
			if completed == course {
				fmt.Println("hey, we found a clash:\n", completed)
			}
		}
	}
	//make slice
	courses := make([]string, 10, 20)
	fmt.Printf("len %d, cap %d\n", len(courses), cap(courses))
}
