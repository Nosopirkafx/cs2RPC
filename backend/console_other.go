//go:build !windows

package main

func enableConsoleColors() bool { return false }
func consoleWidth() int         { return 80 }
