package telemetry

import (
	"fmt"
	"github.com/Fuelvine/f1-telemetry/pkg/packets"
	f1 "github.com/Fuelvine/f1-telemetry/pkg/telemetry"
	"github.com/Fuelvine/kiwi/backend/app"
	"github.com/iMeisa/errortrace"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func setup(app *app.App, tm *Telemetry) (*f1.Client, errortrace.ErrorTrace) {

	client, err := f1.NewClient()
	if err != nil {
		trace := errortrace.NewTrace(err)
		return client, trace
	}

	client.OnEventPacket(func(packet *packets.PacketEventData) {
		tm.TelemetryCode = fmt.Sprintf("%s", packet.EventCodeString())
		runtime.EventsEmit(app.Ctx, "event", packet.EventCodeString())
	})

	client.OnCarTelemetryPacket(func(packet *packets.PacketCarTelemetryData) {
		speed := packet.CarTelemetryData[0].Speed
		tm.SpeedData = append(tm.SpeedData, speed)
		runtime.EventsEmit(app.Ctx, "car", speed)
	})

	return client, errortrace.NilTrace()
}
