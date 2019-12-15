package main

import "fmt"

func main() {
	if firstRank, secondRank := "678", "39"; firstRank < secondRank {
		fmt.Println("First is doing better")
	} else if firstRank == secondRank {
		fmt.Println("First is equals" +
			" to second")
	} else {
		fmt.Println("Second is doing better")
	}

	// switch <simple expr>; <condition>{
	// case <expr>: <code>
	// case <expr>: <code>
	// default: <code>
	// }
	msg := ""
	switch "linux" {
	case "docker":
		msg = "Here Docker courses.."
		fmt.Println(msg)
	case "linux":
		msg = "Here Linux courses.."
		fmt.Println(msg)
		fallthrough
	case "windows":
		msg = "Here Windows courses.."
		fmt.Println(msg)
	default:
		msg = "Sorry, didnt find a match.."
		fmt.Println(msg)
	}
	// fmt.Println(msg)
}
