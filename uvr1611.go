package main
import (
    "gouvr/uvr"
    "fmt"
)

func main() {
    packetReceiver := uvr.NewPacketReceiver()
    packetDecoder := uvr.NewPacketDecoder(packetReceiver)
    byteDecoder     := uvr.NewByteDecoder(packetDecoder, uvr.NewTimeout(488.0, 0.2))
    syncDecoder     := uvr.NewSyncDecoder(byteDecoder, byteDecoder, uvr.NewTimeout(488.0*2, 0.3))
    interrupt := uvr.NewEdgeInterrupt(syncDecoder)
    replayer := uvr.NewReplayer(interrupt)
    
    packetReceiver.RegisterCallback(func(packet uvr.Packet) {
        println("Received packet")
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