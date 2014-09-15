package main
import (
    "gouvr/uvr"
    "fmt"
    "math/big"
    "github.com/kidoman/embd"
    _"github.com/kidoman/embd/host/bbb"
)

func Init(file string) (embd.DigitalPin, error) {
    embd.InitGPIO()
    pin, pin_err := embd.NewDigitalPin(file)
    if pin_err != nil {
		fmt.Printf("Error opening pin! %s\n", pin_err)
		return nil, pin_err
	}
    
	// clean up on exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Println("Closing pin and terminating program.")
            pin.Close()
            embd.CloseGPIO()
			os.Exit(0)
		}
	}()

    pin.SetDirection(embd.In)
    
    return pin, nil
}

func main() {
    pin, err := Init("P8_07")
    if err != nil {
        return
    }
    
    defer pin.Close()
    defer embd.CloseGPIO()
    
    packetReceiver  := uvr.NewPacketReceiver()
    packetDecoder   := uvr.NewPacketDecoder(packetReceiver)
    byteDecoder     := uvr.NewByteDecoder(packetDecoder, uvr.NewTimeout(488.0, 0.2))
    syncDecoder     := uvr.NewSyncDecoder(byteDecoder, byteDecoder, uvr.NewTimeout(488.0*2, 0.3))
    signal          := uvr.Signal(syncDecoder)
    
    packetReceiver.RegisterCallback(func(packet uvr.Packet) {
        packet.Log()
        syncDecoder.Reset()
        byteDecoder.Reset()
        packetDecoder.Reset()
    })
    
    err = pin.Watch(embd.EdgeBoth, func(pin embd.DigitalPin) {
        value, read_err := pin.Read()
        if read_err != nil {
            fmt.Println(read_err)
        } else {
            signal.Consume(big.Word(value))
        }
    })
    
    if err != nil {
	    fmt.Println("Could not watch pin.", err)
    }
    
    select {
        
    }
}