package main

import (
	"chris/cmarket/app"
)

func init() {
	//
}

func main() {
	err := app.SetupAndRunApp()

	if err != nil {
		panic(err)
	}
}
