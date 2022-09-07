package main

import (
	"embed"
	"fmt"
	"github.com/Fuelvine/kiwi/backend/app"
	"github.com/Fuelvine/kiwi/backend/server"
	"github.com/Fuelvine/kiwi/backend/telemetry"
	"github.com/iMeisa/errortrace"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	a := app.NewApp()
	tm, trace := telemetry.Start(a)
	if trace.HasError() {
		trace.Read()
		panic(trace.ErrorString())
	}

	// Start web server
	trace = server.StartServer()
	if trace.HasError() {
		trace.Read()
		panic(trace.ErrorString())
	}

	// Create application with options
	fmt.Println("Starting Wails application")
	err := wails.Run(&options.App{
		Title:     "Kiwi",
		Width:     1000,
		Height:    720,
		Assets:    assets,
		OnStartup: a.Startup,
		Bind: []interface{}{ // Binds struct functions to JS
			a,
			tm,
		},
	})

	if err != nil {
		trace = errortrace.NewTrace(err)
		trace.Read()
	}
}
