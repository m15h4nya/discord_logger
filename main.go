package main

func main() {
	ch := make(chan int, 1)
	go createServer()
	go createSession()
	<-ch
}
