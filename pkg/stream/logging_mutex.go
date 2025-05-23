package stream

import (
	"log"
	"runtime"
	"sync"
	"time"
)

// LoggingMutex is a wrapper around sync.Mutex with logging on Lock and Unlock
type LoggingMutex struct {
	//mu       delock.Mutex
	mu       sync.Mutex
	skipLogs bool
	clientID string
	lockID   int
}

// Lock logs lock operation with caller information and acquires the mutex
func (l *LoggingMutex) Lock() {
	if !l.skipLogs {
		pc, file, line, ok := runtime.Caller(1) // Get caller information
		if ok {
			caller := runtime.FuncForPC(pc).Name()
			log.Printf("[Locking] Mutex locking for '%s' at %s:%d by %s at %v", l.clientID, file, line, caller, time.Now())
			defer log.Printf("[Locked] Mutex locked for '%s' at %s:%d by %s at %v", l.clientID, file, line, caller, time.Now())
		} else {
			log.Printf("[Locking] Mutex locking for '%s' (caller info unavailable) at %v", l.clientID, time.Now())
		}
	}
	l.mu.Lock()
	//var err error
	//l.lockID, err = l.mu.Lock()
	//if err != nil {
	//	panic(err)
	//}
}

// Unlock logs unlock operation with caller information and releases the mutex
func (l *LoggingMutex) Unlock() {
	if !l.skipLogs {
		pc, file, line, ok := runtime.Caller(1) // Get caller information
		if ok {
			caller := runtime.FuncForPC(pc).Name()
			log.Printf("[Unlocking] Mutex unlocking for '%s' at %s:%d by %s at %v", l.clientID, file, line, caller, time.Now())
			log.Printf("[Unlocked] Mutex unlocked for '%s' at %s:%d by %s at %v", l.clientID, file, line, caller, time.Now())
		} else {
			log.Printf("[Unlocking] Mutex unlocked for '%s' (caller info unavailable) at %v", l.clientID, time.Now())
		}
	}
	l.mu.Unlock()
	//
	//l.mu.Unlock(l.lockID)
}
