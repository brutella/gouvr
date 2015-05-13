package uvr1611

import (
	"github.com/brutella/gouvr/uvr"
	"testing"
)

func TestIntegration(t *testing.T) {
	log := "../../logs/integration.log"

	packetReceiver := NewPacketReceiver()
	packetDecoder := NewPacketDecoder(packetReceiver)
	byteDecoder := uvr.NewByteDecoder(packetDecoder, uvr.NewTimeout(SignalFrequency, 0.2))
	syncDecoder := NewSyncDecoder(byteDecoder, byteDecoder, uvr.NewTimeout(SignalFrequency*2, 0.3))
	interrupt := uvr.NewEdgeInterrupt(syncDecoder)
	replayer := uvr.NewReplayer(interrupt)

	var p Packet
	packetReceiver.RegisterCallback(func(packet Packet) {
		p = packet
		syncDecoder.Reset()
		byteDecoder.Reset()
		packetDecoder.Reset()
	})
	if err := replayer.Replay(log); err != nil {
		t.Fatal(err)
	}

	it, v := DecodeInputValue(p.Input1)
	if is, want := it, InputTypeTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(8.5); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input2)
	if is, want := it, InputTypeTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(55.6); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input3)
	if is, want := it, InputTypeTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(45.5); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input4)
	if is, want := it, InputTypeDigital; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(0.0); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input5)
	if is, want := it, InputTypeDigital; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(0.0); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input6)
	if is, want := it, InputTypeDigital; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(0.0); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input7)
	if is, want := it, InputTypeTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(22.4); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input8)
	if is, want := it, InputTypeTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(1.5); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input9)
	if is, want := it, InputTypeTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(73.9); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input10)
	if is, want := it, InputTypeTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(46.7); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input11)
	if is, want := it, InputTypeTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(37.9); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input12)
	if is, want := it, InputTypeTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(2.6); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input13)
	m := RoomTemperatureModeFromValue(p.Input13)
	if is, want := it, InputTypeRoomTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(21.0); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := m, RoomTemperatureModeAutomatic; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input14)
	m = RoomTemperatureModeFromValue(p.Input14)
	if is, want := it, InputTypeRoomTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(20.9); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := m, RoomTemperatureModeAutomatic; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input15)
	if is, want := it, InputTypeTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(17.5); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	it, v = DecodeInputValue(p.Input16)
	if is, want := it, InputTypeVolumeFlow; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := v, float32(0.0); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	outlets := OutletsFromValue(p.Outgoing)
	if is, want := len(outlets), 13; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[0].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[1].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[2].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[3].Enabled, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[4].Enabled, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[5].Enabled, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[6].Enabled, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[7].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[8].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[9].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[10].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[11].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := outlets[12].Enabled, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	h1, h2 := AreHeatMetersEnabled(p.HeatRegister)
	if is, want := h1, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := h2, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

}
