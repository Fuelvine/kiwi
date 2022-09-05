package main

import (
	"fmt"
	"github.com/Fuelvine/f1-telemetry/pkg/packets"
)

type Telemetry struct {
	TelemetryCode string
	packets.CarTelemetryData
}

func NewTelemetry() *Telemetry {
	return &Telemetry{}
}

func (t *Telemetry) GetCode() string {
	return t.TelemetryCode
}

func (t *Telemetry) GetSpeed() uint16 {
	fmt.Print(".")
	return t.CarTelemetryData.Speed
}
