package uvr

// Logs and hands over bit stream to next bit consumer
type handover struct {
    consumer BitConsumer
    logger Logger
}

func NewHandover(consumer BitConsumer, l Logger) *handover {
    h := &handover{consumer: consumer}
    h.logger = l
        
    return h
}

func (h *handover) Consume(bit Bit) error {
    h.logger.Log(LogString(bit) + "\n")            
    h.consumer.Consume(bit)
    return nil
}

func (h *handover) Reset() {
}