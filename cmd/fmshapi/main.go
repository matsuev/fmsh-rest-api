package main

func main() {
	println("FMSH REST API Server")

	app := NewApplication("gin")

	app.Start()

	app.Shutdown()
}
