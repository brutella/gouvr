package uvr1611

type PacketCallback func(packet Packet)
// Packet receiver
type packetReceiver struct {
    PacketConsumer
    Packet Packet
    
    callback PacketCallback
}

func (p *packetReceiver) Consume(packet Packet) error {
    p.Packet = packet
    
    if p.callback != nil {
        p.callback(p.Packet)
    }
    
    return nil
}

func NewPacketReceiver() *packetReceiver {
    return &packetReceiver{}
}

func (p *packetReceiver) RegisterCallback(callback PacketCallback) {
    p.callback = callback
}