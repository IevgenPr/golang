package select_

// This is sample from Rob Pike preso https://talks.golang.org/2012/waza.slide#35
import (
	"fmt"
	"math/rand"
	"time"
)

func select_() {
	timerChan := make(chan time.Time)
	fmt.Println("Starting go routine at", time.Now())
	rand.Seed(time.Now().UnixNano()) // seed with current time to get rand working
	someInt := rand.Intn(10)
	deltaT := time.Duration(someInt) * time.Second
	go func() {
		time.Sleep(deltaT)
		timerChan <- time.Now()
	}() // last parenthesis are important for closure to work

	fmt.Println("Started go routine at", time.Now())
	completedAt := <-timerChan
	fmt.Println("Completed go routine at", completedAt, "waited", deltaT, someInt)

	// waits for messages from either channels, if no messages - defaults
	// select {
	// case v := <-ch1:
	// 	fmt.Println("ch1 sends ", v)
	// case v := <-ch1:
	// 	fmt.Println("ch2 sends ", v)
	// default: // fall through, othewise wait
	// 	fmt.Println("default ")

	// }

}
