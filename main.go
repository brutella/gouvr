package main

import(
    _"gouvr/uvr"
    "log"
    "github.com/kidoman/embd"
    _"github.com/kidoman/embd/host/bbb"
)

func main() {
    embd.InitGPIO()
    defer embd.CloseGPIO()
    pin, err := embd.NewDigitalPin(10)
    if err != nil {
        log.Println("Could not access digital pin.", err)
        return
    }
    
    pin.Read()
}