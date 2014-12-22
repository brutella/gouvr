package uvr1611

type PacketCallback func(packet Packet)

type packetReceiver struct {
    PacketConsumer
    Packet Packet
    
    callback PacketCallback
}

// NewPacketReceiver returns a packet receiver, which implements the PacketConsumer interface.
func NewPacketReceiver() *packetReceiver {
    return &packetReceiver{}
}

// RegisterCallback registers a function which is called when a new packet is consumed.
func (p *packetReceiver) RegisterCallback(callback PacketCallback) {
    p.callback = callback
}

func (p *packetReceiver) Consume(packet Packet) error {
    p.Packet = packet
    
    if p.callback != nil {
        p.callback(p.Packet)
    }
    
    return nil
}