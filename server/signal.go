package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func signalHandler() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)

	go func() {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT:
			fmt.Printf("%s received. Exit\n", sig.String())
			os.Exit(1)
		default:
			fmt.Println("Unhandled signal")
		}
	}()
}
