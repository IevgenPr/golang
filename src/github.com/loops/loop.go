package main

import "fmt"

func main() {
	// slice - unordered list
	coursesInProg := []string{"Docker deep dive", "Docker Clustering",
		"Docker and Kubernetes"}

	for _, course := range coursesInProg {
		fmt.Println(course)
	}
}
