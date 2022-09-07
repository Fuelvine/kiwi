package telemetry

import (
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

	client.OnMotionPacket(func(packet *packets.PacketMotionData) {

		frame := GetFrame(packet.Header.FrameIdentifier)
		frame.PacketMotionData = *packet
		Frames[packet.Header.FrameIdentifier] = frame
	})

	client.OnSessionPacket(func(packet *packets.PacketSessionData) {

		frame := GetFrame(packet.Header.FrameIdentifier)
		frame.PacketSessionData = *packet
		Frames[packet.Header.FrameIdentifier] = frame
	})

	client.OnLapPacket(func(packet *packets.PacketLapData) {

		frame := GetFrame(packet.Header.FrameIdentifier)
		frame.PacketLapData = *packet
		Frames[packet.Header.FrameIdentifier] = frame
	})

	client.OnEventPacket(func(packet *packets.PacketEventData) {

		frame := GetFrame(packet.Header.FrameIdentifier)
		frame.PacketEventData = *packet
		Frames[packet.Header.FrameIdentifier] = frame
		runtime.EventsEmit(app.Ctx, "event", packet.EventCodeString())
	})

	client.OnParticipantsPacket(func(packet *packets.PacketParticipantsData) {

		frame := GetFrame(packet.Header.FrameIdentifier)
		frame.PacketParticipantsData = *packet
		Frames[packet.Header.FrameIdentifier] = frame
	})

	client.OnCarSetupPacket(func(packet *packets.PacketCarSetupData) {

		frame := GetFrame(packet.Header.FrameIdentifier)
		frame.PacketCarSetupData = *packet
		Frames[packet.Header.FrameIdentifier] = frame
	})

	client.OnCarTelemetryPacket(func(packet *packets.PacketCarTelemetryData) {
		//AddToFrames(packet.Header.FrameIdentifier, packet)

		frame := GetFrame(packet.Header.FrameIdentifier)
		frame.PacketCarTelemetryData = *packet
		Frames[packet.Header.FrameIdentifier] = frame

		speed := packet.CarTelemetryData[0].Speed
		runtime.EventsEmit(app.Ctx, "car", speed)
	})

	client.OnCarStatusPacket(func(packet *packets.PacketCarStatusData) {

		frame := GetFrame(packet.Header.FrameIdentifier)
		frame.PacketCarStatusData = *packet
		Frames[packet.Header.FrameIdentifier] = frame
	})

	client.OnFinalClassificationPacket(func(packet *packets.PacketFinalClassificationData) {

		frame := GetFrame(packet.Header.FrameIdentifier)
		frame.PacketFinalClassificationData = *packet
		Frames[packet.Header.FrameIdentifier] = frame
	})

	client.OnLobbyInfoPacket(func(packet *packets.PacketLobbyInfoData) {

		frame := GetFrame(packet.Header.FrameIdentifier)
		frame.PacketLobbyInfoData = *packet
		Frames[packet.Header.FrameIdentifier] = frame
	})

	client.OnCarDamagePacket(func(packet *packets.PacketCarDamageData) {

		frame := GetFrame(packet.Header.FrameIdentifier)
		frame.PacketCarDamageData = *packet
		Frames[packet.Header.FrameIdentifier] = frame
	})

	client.OnSessionHistoryPacket(func(packet *packets.PacketSessionHistoryData) {

		frame := GetFrame(packet.Header.FrameIdentifier)
		frame.PacketSessionHistoryData = *packet
		Frames[packet.Header.FrameIdentifier] = frame
	})

	return client, errortrace.NilTrace()
}
