package main
import (
    "gouvr/uvr"
    "fmt"
)

const LOG_FILE = "./logs/2014-09-14_interrupt.log"
func main() {
    packetReceiver := uvr.NewUVR1611PacketReceiver()
    packetDecoder := uvr.NewUVR1611PacketDecoder(packetReceiver)
    byteDecoder     := uvr.NewByteDecoder(packetDecoder, uvr.NewTimeout(488.0, 0.2))
    syncDecoder     := uvr.NewUVR1611SyncDecoder(byteDecoder, byteDecoder, uvr.NewTimeout(488.0*2, 0.3))
    interrupt := uvr.NewEdgeInterrupt(syncDecoder)
    replayer := uvr.NewReplayer(interrupt)
    
    packetReceiver.RegisterCallback(func(packet uvr.UVR1611Packet) {
        packet.Log()
        syncDecoder.Reset()
        byteDecoder.Reset()
    })
    err := replayer.Replay(LOG_FILE)
    if err != nil {
        fmt.Println("Could not replay file.", err)
    }
    
    // packetReceiver.Packet.Log()
}