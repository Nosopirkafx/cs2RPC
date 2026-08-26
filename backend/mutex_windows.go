//go:build windows

package main

import "golang.org/x/sys/windows"

const mutexName = "Local\\FACEIT_RPC_SingleInstance"

var mutexHandle windows.Handle

func alreadyRunning() bool {
	h, err := windows.CreateMutex(nil, false, windows.StringToUTF16Ptr(mutexName))
	if err != nil {
		return false
	}
	if windows.GetLastError() == windows.ERROR_ALREADY_EXISTS {
		_ = windows.CloseHandle(h)
		return true
	}
	mutexHandle = h
	return false
}
