package uvr1611

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/brutella/gouvr/uvr"
)

func TestIntegration(test *testing.T) {
    log := "../../logs/integration.log"
    
    packetReceiver  := NewPacketReceiver()
    packetDecoder   := NewPacketDecoder(packetReceiver)
    byteDecoder     := uvr.NewByteDecoder(packetDecoder, uvr.NewTimeout(SignalFrequency, 0.2))
    syncDecoder     := NewSyncDecoder(byteDecoder, byteDecoder, uvr.NewTimeout(SignalFrequency*2, 0.3))
    interrupt       := uvr.NewEdgeInterrupt(syncDecoder)
    replayer        := uvr.NewReplayer(interrupt)
    
    var p Packet
    packetReceiver.RegisterCallback(func(packet Packet) {
        p = packet
        syncDecoder.Reset()
        byteDecoder.Reset()
        packetDecoder.Reset()
    })
    err := replayer.Replay(log)
    assert.Nil(test, err)
    assert.NotNil(test, p)
    
    t, v := DecodeInputValue(p.Input1)
    assert.Equal(test, t, InputTypeTemperature)
    assert.Equal(test, v, 8.5)
    
    t, v = DecodeInputValue(p.Input2)
    assert.Equal(test, t, InputTypeTemperature)
    assert.Equal(test, v, 55.6)
    
    t, v = DecodeInputValue(p.Input3)
    assert.Equal(test, t, InputTypeTemperature)
    assert.Equal(test, v, 45.5)
    
    t, v = DecodeInputValue(p.Input4)
    assert.Equal(test, t, InputTypeDigital)
    assert.Equal(test, v, 0.0)
    
    t, v = DecodeInputValue(p.Input5)
    assert.Equal(test, t, InputTypeDigital)
    assert.Equal(test, v, 0.0)
    
    t, v = DecodeInputValue(p.Input6)
    assert.Equal(test, t, InputTypeDigital)
    assert.Equal(test, v, 0.0)
    
    t, v = DecodeInputValue(p.Input7)
    assert.Equal(test, t, InputTypeTemperature)
    assert.Equal(test, v, 22.4)
    
    t, v = DecodeInputValue(p.Input8)
    assert.Equal(test, t, InputTypeTemperature)
    assert.Equal(test, v, 1.5)
    
    t, v = DecodeInputValue(p.Input9)
    assert.Equal(test, t, InputTypeTemperature)
    assert.Equal(test, v, 73.9)
    
    t, v = DecodeInputValue(p.Input10)
    assert.Equal(test, t, InputTypeTemperature)
    assert.Equal(test, v, 46.7)
    
    t, v = DecodeInputValue(p.Input11)
    assert.Equal(test, t, InputTypeTemperature)
    assert.Equal(test, v, 37.9)
    
    t, v = DecodeInputValue(p.Input12)
    assert.Equal(test, t, InputTypeTemperature)
    assert.Equal(test, v, 2.6)
    
    t, v = DecodeInputValue(p.Input13)
    m := RoomTemperatureModeFromValue(p.Input13)
    assert.Equal(test, t, InputTypeRoomTemperature)
    assert.Equal(test, v, 21.0)
    assert.Equal(test, m, RoomTemperatureModeAutomatic)
    
    t, v = DecodeInputValue(p.Input14)
    m = RoomTemperatureModeFromValue(p.Input14)
    assert.Equal(test, t, InputTypeRoomTemperature)
    assert.Equal(test, v, 20.9)
    assert.Equal(test, m, RoomTemperatureModeAutomatic)
    
    t, v = DecodeInputValue(p.Input15)
    assert.Equal(test, t, InputTypeTemperature)
    assert.Equal(test, v, 17.5)
    
    t, v = DecodeInputValue(p.Input16)
    assert.Equal(test, t, InputTypeVolumeFlow)
    assert.Equal(test, v, 0.0)
    
    outlets := OutletsFromValue(p.Outgoing)
    assert.Equal(test, len(outlets), 13)
    assert.False(test, outlets[0].Enabled)
    assert.False(test, outlets[1].Enabled)
    assert.False(test, outlets[2].Enabled)
    assert.True(test, outlets[3].Enabled)
    assert.True(test, outlets[4].Enabled)
    assert.True(test, outlets[5].Enabled)
    assert.True(test, outlets[6].Enabled)
    assert.False(test, outlets[7].Enabled)
    assert.False(test, outlets[8].Enabled)
    assert.False(test, outlets[9].Enabled)
    assert.False(test, outlets[10].Enabled)
    assert.False(test, outlets[11].Enabled)
    assert.False(test, outlets[12].Enabled)
    
    h1, h2 := AreHeatMetersEnabled(p.HeatRegister)
    assert.True(test, h1)
    assert.False(test, h2)
    
}
