package main

import (
	"fmt"
	"reflect"
)

func main() {
	myCourses := make([]int, 10, 20) // make(type, len, capacity)
	//can drop capacity
	tCourses := []string{"COurse 1", "course2", "course3"}
	// tCourses.appen
	fmt.Println(len(myCourses), cap(myCourses), myCourses, tCourses)
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySlice[0] = 22
	fmt.Println(mySlice)
	sliceOfSlice := mySlice[2:5]
	fmt.Println(sliceOfSlice)

	mSl := make([]int, 1, 4)
	fmt.Println(len(mSl), cap(mSl))
	for i := 1; i < 17; i++ {
		mSl = append(mSl, i)
		fmt.Println(i, len(mSl), cap(mSl))

	}
	newSl := []int{44, 55, 66}
	fmt.Println(reflect.TypeOf(newSl), reflect.TypeOf(mSl))
	newSl = append(newSl, mSl...)
	fmt.Println(newSl)
}
