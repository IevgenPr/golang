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
	switch "docker" {
	case "docker":
		msg = "Here Docker courses.."
	case "linux":
		msg = "Here Linux courses.."
	case "windows":
		msg = "Here Windows courses.."
	default:
		msg = "Sorry, didnt find a match.."
	}
	fmt.Println(msg)
}
