package main

import "sync"

func main() {
	wg := sync.WaitGroup{}
	server := HTTPService{}
	server.InitService()
	wg.Add(1)
	go server.createServer()
	wg.Wait()
}
