package main

func main() {
	ctx := NewRestServerContext()
	server := NewRestServer(ctx)
	server.InitializeRoutes()
	server.Start()
}
