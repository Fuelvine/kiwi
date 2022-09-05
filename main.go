package main

import (
	"embed"
	"fmt"
	"github.com/Fuelvine/f1-telemetry/pkg/packets"
	"github.com/Fuelvine/f1-telemetry/pkg/telemetry"
	"github.com/iMeisa/errortrace"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	client, err := telemetry.NewClient()
	if err != nil {
		trace := errortrace.NewTrace(err)
		trace.Read()
		panic(trace.ErrorString())
	}

	client.OnEventPacket(func(packet *packets.PacketEventData) {
		fmt.Println("Code: ", packet.EventStringCode)
	})

	fmt.Println("Telemetry client started")
	go client.Run()

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	fmt.Println("Starting Wails application")
	err = wails.Run(&options.App{
		Title:            "kiwi",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
