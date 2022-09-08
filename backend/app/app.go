package app

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx

	runtime.WindowSetMinSize(ctx, 800, 500)
	go a.watchWindowSize()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) watchWindowSize() {
	prevWidth, prevHeight := runtime.WindowGetSize(a.Ctx)
	currentWidth, currentHeight := prevWidth, prevHeight
	for {
		currentWidth, currentHeight = runtime.WindowGetSize(a.Ctx)
		if currentWidth != prevWidth || currentHeight != prevHeight {
			prevWidth, prevHeight = currentWidth, currentHeight
			runtime.EventsEmit(a.Ctx, "windowSizeChanged", currentWidth, currentHeight)
		}
	}
}
