package main

import(
    "gouvr/uvr"
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
    }    
    _, err = uvr.New(pin, uvr.DeviceTypeUVR1611)
    if err != nil {
        log.Println("Could create UVR 1611", err)
    }
}