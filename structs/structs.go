package main

import "fmt"

type courseMeta struct {
	Author string
	Level  string
	Rating float64
}

var (
	// DockerDeepDive Some nice variable
	DockerDeepDive = courseMeta{
		Author: "Nice Fellow",
		Level:  "Intermediate",
		Rating: 5,
	}
)

func main() {
	// two declaration options
	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta) //this gives a pointer with zeroed flds

	fmt.Println(DockerDeepDive)
	DockerDeepDive.Rating = 1
	fmt.Printf("Course author: %s\n", DockerDeepDive.Author)
	fmt.Printf("Course raiting: %.1f\n", DockerDeepDive.Rating)
}
