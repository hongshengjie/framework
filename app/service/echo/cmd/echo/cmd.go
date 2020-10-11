package main

import (
	"framework/app/service/echo/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	server.Start()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sign := <-c
		switch sign {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			server.Close()
			time.Sleep(time.Second)
			os.Exit(1)
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
