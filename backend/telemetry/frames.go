package telemetry

import (
	"encoding/json"
	"github.com/Fuelvine/f1-telemetry/pkg/packets"
	"io/ioutil"
)

type Frame struct {
	// Frame number
	Number uint32 `json:"number"`

	// Packets
	packets.PacketCarDamageData
	packets.PacketCarSetupData
	packets.PacketCarStatusData
	packets.PacketCarTelemetryData
	packets.PacketEventData
	packets.PacketFinalClassificationData
	packets.PacketLapData
	packets.PacketLobbyInfoData
	packets.PacketMotionData
	packets.PacketParticipantsData
	packets.PacketSessionData
	packets.PacketSessionHistoryData
}

var Frames = make(map[uint32]Frame)

func GetFrame(frameIdentifier uint32) Frame {

	// Check if frame exists
	frame, ok := Frames[frameIdentifier]

	if !ok {
		return Frame{}
	}

	return frame
}

func GetLastFrame() Frame {
	return Frames[uint32(len(Frames)-1)]
}

func (t *Telemetry) FramesJSON() {
	// Write frames map to JSON
	framesJSON, _ := json.MarshalIndent(Frames, "", "  ")

	// Write JSON to file
	_ = ioutil.WriteFile("frames.json", framesJSON, 0644)
}
