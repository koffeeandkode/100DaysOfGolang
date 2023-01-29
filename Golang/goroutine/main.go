/*
1. ctrl+space for suggestions
2. simple greeter to print 2 string
3. add go to first func-call
- doesn't print it
- why? fireup thread but you didn't wait. main exits before that.
4. So? Hack?
- time.Sleep in called-func
- sync package. better than time.sleep
5. PROD-Scenarios
- we talk to lot of db. 1) read-replicas 2) read-write servers
- we talk to lot of micro-services
- 3 different apis
6. getStautsCode func()
- adding will trigger multiple go-routines || doesn't reply || we are not waiting
7. WaitGroup
- variable vs pointer
- Add(), Done(), Wait()
- Add() - will wait for the go-routine
- Done() - your job to make it done
- Wait() - optional || doesn't allow main() to finish
8. Add these to getStautsCode
9. Adding signals
- append the signals
- issue with goroutines. managed by goRuntime.
- what if 5 different goroutines. they try to write to a single memory.
- mutex helps here
10. Mutex
- gives lock over memory. doesn't give access to other
- readWrite Mutex. (reading allowed. when writing, it kicks out current one reading)
- especially in databases

. Questions
- what if I give incorrect number in the wait-group
-
*/
package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup //usually pointers
var mut sync.Mutex    //usually pointers. since we need to pass around.

func main() {
	// go greeter("HELLO")
	// greeter("WORLD")
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
	}
	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	// for i := 0; i < 6; i++ {
// 	// 	time.Sleep(1 * time.Second)
// 	// 	fmt.Printf("%s-%d \n", s, i)
// 	// }
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("ERROR \n")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d STATUS \n", res.StatusCode)
	}
}

// # Personal GitHub account:
// Host home-github.com
//   HostName github.com
//   User git
//   IdentityFile ~/.ssh/id_rsa.home
//   PreferredAuthentications publickey
//   PasswordAuthentication no
//   IdentitiesOnly yes

//   git@home-github.com:sroy8091/...

// # Work GitHub account:
// Host github.com
//   HostName github.com
//   User git
//   IdentityFile ~/.ssh/id_rsa
//   PreferredAuthentications publickey
//   PasswordAuthentication no
//   IdentitiesOnly yes
