package main

type Telemetry struct {
	TelemetryCode string
}

func NewTelemetry() *Telemetry {
	return &Telemetry{}
}

func (t *Telemetry) UpdateCode() string {
	return t.TelemetryCode
}
