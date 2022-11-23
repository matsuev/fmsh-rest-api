package main

// main function
func main() {
	println("FMSH REST API Server")

	app := NewApplication("gin")

	app.Start()

	app.Shutdown()
}
