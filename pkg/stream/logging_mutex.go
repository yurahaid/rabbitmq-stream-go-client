package stream

import (
	"log"
	"runtime"
	"sync"
	"time"
)

// LoggingMutex is a wrapper around sync.Mutex with logging on Lock and Unlock
type LoggingMutex struct {
	mu       sync.Mutex
	clientID string
}

// Lock logs lock operation with caller information and acquires the mutex
func (l *LoggingMutex) Lock() {
	pc, file, line, ok := runtime.Caller(1) // Get caller information
	if ok {
		caller := runtime.FuncForPC(pc).Name()
		log.Printf("[Locking] Mutex locked for clientID '%s' at %s:%d by %s at %v", l.clientID, file, line, caller, time.Now())
	} else {
		log.Printf("[Locking] Mutex locked for clientID '%s' (caller info unavailable) at %v", l.clientID, time.Now())
	}
	l.mu.Lock()
}

// Unlock logs unlock operation with caller information and releases the mutex
func (l *LoggingMutex) Unlock() {
	pc, file, line, ok := runtime.Caller(1) // Get caller information
	if ok {
		caller := runtime.FuncForPC(pc).Name()
		log.Printf("[Unlocking] Mutex unlocked for clientID '%s' at %s:%d by %s at %v", l.clientID, file, line, caller, time.Now())
	} else {
		log.Printf("[Unlocking] Mutex unlocked for clientID '%s' (caller info unavailable) at %v", l.clientID, time.Now())
	}
	l.mu.Unlock()
}
