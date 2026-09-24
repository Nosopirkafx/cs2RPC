//go:build windows

package main

import (
	"errors"

	"golang.org/x/sys/windows"
)

const mutexName = "Local\\FACEIT_RPC_SingleInstance"

var mutexHandle windows.Handle

func alreadyRunning() bool {
	h, err := windows.CreateMutex(nil, false, windows.StringToUTF16Ptr(mutexName))
	if errors.Is(err, windows.ERROR_ALREADY_EXISTS) {
		_ = windows.CloseHandle(h)
		return true
	}
	if err != nil {
		return false
	}
	mutexHandle = h
	return false
}
