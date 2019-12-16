package main

import "fmt"

func main() {
	//sliceSample = make([]int)        //declare empty slice slice has len and capacity
	prLeague := make(map[string]int) //declare map
	prLeague["Su"] = 6
	prLeague["Ne"] = 4
	//map[Ne:4 Su:6]
	fmt.Println(prLeague)
	//composite literal form
	recentHeadToHead := map[string]int{
		"Su": 5,
		"Ne": 3,
	}
	fmt.Printf("\nLeague Titles: %v\nRecent Head to Head: %v\n",
		prLeague, recentHeadToHead)

	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}

	for k, v := range testMap {
		fmt.Printf("Key %v, Value %v\n", k, v)
	}
	testMap["A"] = 100
	testMap["F"] = 19889
	fmt.Println(testMap["A"])
	printMap(&testMap)
	delete(testMap, "F")
	printMap(&testMap)

}

func printMap(m *map[string]int) {
	fmt.Println("----")
	for k, v := range *m {
		fmt.Printf("Key %v, Value %v\n", k, v)
	}
}
