/*
1. Need for channels
- multiple GR access same memory. solved using mutex
- can handoff some information on-the go
- 1 GR is not aware of other GR
- way in which multiple GR can talk to each other
2. Channels
- pipeline through which multiple channels interact
- arrow to the left always
3. Errors
=> (Classic) fatal error: all goroutines are asleep - deadlock!
- only if somebody is listening to me, can't pass value int the channel.
- can't utilise the channel if nobody is listening
=> send on closed channel
- listening is allowed
- guarantee of getting a zero? (is it a value or reading from a closed channel)
-

4. Code
- 2 GR. 1st to read channel value, 2nd to write into it
- add 2 value. 1 print statement => returns errors. needs listeners
5. Buffered Channel
- myCh := make(chan int, 1)
- doesn't return error in case of 4.2 as it takes only 1 value
6. Close Channel
-
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	myCh := make(chan int, 2)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {

		myCh <- 0
		myCh <- 6
		close(myCh)
		wg.Done()

	}(myCh, wg)
	wg.Wait()

	// myCh <- 5
	// fmt.Println(<-myCh)

}
