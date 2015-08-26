
package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.Tick(1 * time.Second)
	boom := time.After(10 * time.Second)
	counter := 10

	fmt.Println(counter)
	counter--

	for {
		select {
		case <- tick:
			fmt.Println(counter)
			counter--
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(250 * time.Millisecond)
		}
	}
}
