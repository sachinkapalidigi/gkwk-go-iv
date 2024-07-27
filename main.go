package main

import "gkgk-go-iv/consumer"

func main() {
	poller := consumer.NewPoll()
	poller.Start()
}
