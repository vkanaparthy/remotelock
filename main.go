package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {

	// multiple users obtain a lock (for a resource. movie tickets, db access, remote file access)

	lockManager, err := NewLockManager()
	if err != nil {
		fmt.Println("error creating lock manager")
	}
	lock_key := "my_prime_key"

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func() {

			defer wg.Done()

			owner := fmt.Sprintf("worker-%d", rand.IntN(1000))
			lock_acquired := false
			for !(lock_acquired) {
				//fmt.Printf("user %s is trying to acquire key\n", owner)
				lock_acquired = lockManager.Acquire(lock_key, owner)
				if lock_acquired {
					fmt.Printf("lock acquired by user: %s\n", owner)
					time.Sleep(1)
					//primeNumbersSum(100)
					lockManager.Release(lock_key, owner)
				}
			}

		}()
	}

	wg.Wait()
	fmt.Println("Done!!")
}

type LockManager struct {
	// mutex
	mu *sync.Mutex
	// list of locks held by owners
	locks map[string]string
}

func NewLockManager() (*LockManager, error) {
	var mu = sync.Mutex{}
	return &LockManager{
		mu:    &mu,
		locks: make(map[string]string),
	}, nil
}

// works responsibility to release the key. TODO: implement ttl
/*func doWork(lockMgr *LockManager, n int, lock_key string, worker_name string, wg *sync.WaitGroup) {

	defer wg.Done()
	for { // try to acquire the lock
		lock_acquired := lockMgr.Acquire(lock_key, worker_name)
		if lock_acquired {
			fmt.Printf("User %s acquired the lock: %s\n", worker_name, lock_key)
			// do some work
			primeNumbersSum(n)
			//release the lock
			lockMgr.Release(lock_key, worker_name)
			return
		}
	}
}*/

func primeNumbersSum(n int) {
	// running the for loop from 1 to n
	total_sum := 0
	for i := 2; i <= n; i++ {

		// flag which will confirm that the number is Prime or not
		isPrime := true

		// This for loop is checking that the current number have
		// any divisible other than 1
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		// if the value of the flag is false then the number is not prime
		// and we are not printing it.
		if isPrime {

			total_sum += i
		}
	}
	fmt.Printf("Sum of %d prime numbers is %d", n, total_sum)
}
