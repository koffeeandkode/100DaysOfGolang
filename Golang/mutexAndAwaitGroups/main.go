/*
1. What is race condition
- multiple go-routines
- 2 thread try to write to the same memory
2. Race Tooling
-
3. Anonymous functions
- execute immediately
4. Code
- 3 go-routines
- add waitgroups
-
5. Check race
- go run --race .
- exit status 66
6. Add mutex
- always lock memory
- can use read-write Mutex.
>> mut.Rlock()
>> mut.Unlock()
- always lock when reading. not add at the source.
- if you just use Mutex, not need to put Rlock
- strick lock not necessary for reading
7. RW Mutex
- throws all routines(trying to read) when new thread wants to write.
8. Notes
this s the difference ..Mutex holds a lock for both reads and writes, whereas RwLock treats reads and writes differently, allowing for multiple read locks to be taken in parallel but requiring exclusive access for write locks.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")

	wg := &sync.WaitGroup{}

	mut := &sync.Mutex{}
	rwmut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("ONE R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("RW R")
		rwmut.RLock()
		fmt.Println(score)
		rwmut.RUnlock()
		wg.Done()
	}(wg, rwmut)

	wg.Wait()

}
