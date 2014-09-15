package main

import(
    _"log"
    "github.com/kidoman/embd"
    _"github.com/kidoman/embd/host/bbb"
    "os"
    "os/signal"
    "time"
    "fmt"
)

func main() {
    embd.InitLED()
    embd.InitGPIO()
    
    pin, pin_err := embd.NewDigitalPin("P8_38")
    if pin_err != nil {
		fmt.Printf("Error opening pin! %s\n", pin_err)
		return
	}

    led, led_err := embd.NewLED(3)
    if led_err != nil {
		fmt.Printf("Error accessing led! %s\n", led_err)
		return
	}
    
	// clean up on exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Println("Closing pin and terminating program.")
            led.Off()
            led.Close()
            pin.Close()
            embd.CloseGPIO()
            embd.CloseLED()
			os.Exit(0)
		}
	}()

    pin.SetDirection(embd.Out)
    pin.Write(embd.High)
    
	for {
		fmt.Println("Toggle")
        led.Toggle()
		time.Sleep(2000 * time.Millisecond)
	}
}