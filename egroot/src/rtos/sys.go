package rtos

import "syscall"

// Uptime returns how long system is running (in nanosecond). Time when system
// was in deep sleep state can be included or not.
func Uptime() uint64 {
	return syscall.Uptime()
}

// SleepUntil sleeps task until Uptime() < end.
func SleepUntil(end uint64) {
	sleepUntil(end)
}