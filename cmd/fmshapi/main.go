package main

func main() {
	println("FMSH REST API Server")

	app := NewApplication("http", "gin")

	app.Start()

	app.Shutdown()
}
