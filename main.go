package main

import "github.com/theapemachine/bugger/bugui"

func main() {
	screen := bugui.NewScreen()
	screen.Render().Wait()
}
