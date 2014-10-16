// Writes the GPIO pin values into the pipeline and logs packets to the console
package main

import (
    "fmt"
    "math/big"
    "time"
    "os"
    "os/signal"
    
    "github.com/brutella/gouvr/uvr"
    "github.com/brutella/gouvr/uvr/1611"
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
    
    packetReceiver  := uvr1611.NewPacketReceiver()
    packetDecoder   := uvr1611.NewPacketDecoder(packetReceiver)
    byteDecoder     := uvr.NewByteDecoder(packetDecoder, uvr.NewTimeout(488.0, 0.4))
    syncDecoder     := uvr1611.NewSyncDecoder(byteDecoder, byteDecoder, uvr.NewTimeout(488.0*2, 0.4))
    signal          := uvr.NewSignal(syncDecoder)
    
    packetReceiver.RegisterCallback(func(packet uvr1611.Packet) {
        /*
        (Aussentemperatur)
        (Fussbodenheizung Vorlauf)
        (Buffer Warmwasser - Oben)
        (Buffer Warmwasser - Mitte)
        (Buffer Warmwasser - Unten)
        (Raumtemperatur)
        (Wärmetauscher Sekundär)
        */
        fmt.Println(time.Now().Format(time.Stamp))
        fmt.Println("Zeit:", packet.Timestamp.ToString())
        fmt.Println("Aussentemperatur:", uvr1611.InputValueToString(packet.Input1))
        fmt.Println("Fussbodenheizung Vorlauf:", uvr1611.InputValueToString(packet.Input2))
        fmt.Println("Buffer Warmwasser")
        fmt.Println("   Oben:", uvr1611.InputValueToString(packet.Input3))
        fmt.Println("   Mitte:", uvr1611.InputValueToString(packet.Input4))
        fmt.Println("   Unten:", uvr1611.InputValueToString(packet.Input5))
        fmt.Println("Raumtemperatur:", uvr1611.InputValueToString(packet.Input6))
        fmt.Println("Wärmetauscher Sekundär:", uvr1611.InputValueToString(packet.Input7))
        // packet.Log()
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