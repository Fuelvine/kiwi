package telemetry

import "github.com/Fuelvine/f1-telemetry/pkg/packets"

func (t *Telemetry) GetCode() string {
	return t.TelemetryCode
}

func (t *Telemetry) GetSpeed() uint16 {
	return t.CarTelemetryData.Speed
}

func (t *Telemetry) GetCarTelemetry() packets.CarTelemetryData {
	return t.CarTelemetryData
}
