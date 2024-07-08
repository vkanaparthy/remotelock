package main

import (
	"fmt"
)

func (lockMgr *LockManager) Acquire(lock_key string, owner string) bool {
	lockMgr.mu.Lock()
	result := SetLock(lock_key, owner, lockMgr)
	lockMgr.mu.Unlock()
	return result
}

func (lockMgr *LockManager) Release(lock_key string, owner string) bool {
	lockMgr.mu.Lock()
	result := ClearLock(lock_key, owner, lockMgr)
	lockMgr.mu.Unlock()
	return result
}

func ClearLock(lock_key string, owner string, lockMgr *LockManager) bool {

	current_owner, status := lockMgr.locks[lock_key]
	// lock exists
	if status {
		// does the current user owns the lock
		if current_owner == owner {
			// free the lock

			delete(lockMgr.locks, lock_key)

			fmt.Printf("Lock: %s is released by owner: %s\n", lock_key, current_owner)
			return true
		} else {
			fmt.Printf("Lock: %s is not owned by user: %s\n", lock_key, owner)
			return false
		}
	} else {
		fmt.Printf("Lock: %s is not owned by user: %s\n", lock_key, owner)
		return true
	}
}

func SetLock(lock_key string, owner string, lockMgr *LockManager) bool {

	current_owner, status := lockMgr.locks[lock_key]
	//lock exists
	if status {
		//check who is the owner
		if current_owner != owner { //someone else owns the lock
			//fmt.Printf("Lock : %s is currently owned by: %s\n", lock_key, current_owner)
			return false
		} else {
			// current owner has the lock
			//fmt.Printf("Lock : %s  already owned by: %s\n", lock_key, current_owner)
			return true
		}
	} else {
		//lock is available and is assiged to current owner
		lockMgr.locks[lock_key] = owner
		//fmt.Printf("Lock : %s  is now owned by: %s\n", lock_key, owner)
		return true
	}
}
