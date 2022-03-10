package main

func main() {
	server := HTTPService{}
	server.InitService()
	server.createServer()
}
