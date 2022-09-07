package telemetry

import (
	"fmt"
	"github.com/Fuelvine/kiwi/backend/app"
	"github.com/iMeisa/errortrace"
)

type Telemetry struct {
	TelemetryCode string
}

func NewTelemetry() *Telemetry {
	return &Telemetry{}
}

func Start(app *app.App) (*Telemetry, errortrace.ErrorTrace) {
	tm := NewTelemetry()

	// Create telemetry client
	client, trace := setup(app, tm)
	if trace.HasError() {
		return nil, trace
	}

	fmt.Println("Telemetry telemetry started")
	go client.Run()

	return tm, errortrace.NilTrace()
}
