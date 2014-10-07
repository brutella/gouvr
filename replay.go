package main
import (
    "github.com/brutella/gouvr/uvr"
    "github.com/brutella/gouvr/uvr/1611"
    "fmt"
)

const LOG_FILE = "./logs/2014-09-14_interrupt.log"
func main() {
    packetReceiver := uvr1611.NewPacketReceiver()
    packetDecoder := uvr1611.NewPacketDecoder(packetReceiver)
    byteDecoder     := uvr.NewByteDecoder(packetDecoder, uvr.NewTimeout(488.0, 0.2))
    syncDecoder     := uvr1611.NewSyncDecoder(byteDecoder, byteDecoder, uvr.NewTimeout(488.0*2, 0.3))
    interrupt := uvr.NewEdgeInterrupt(syncDecoder)
    replayer := uvr.NewReplayer(interrupt)
    
    packetReceiver.RegisterCallback(func(packet uvr1611.Packet) {
        packet.Log()
        syncDecoder.Reset()
        byteDecoder.Reset()
        packetDecoder.Reset()
    })
    err := replayer.Replay(LOG_FILE)
    if err != nil {
        fmt.Println("Could not replay file.", err)
    }
    
    packetReceiver.Packet.Log()
}