package main

import(
    "gouvr/uvr"
    "github.com/kidoman/embd"
    _"github.com/kidoman/embd/host/bbb"
    "os"
    "time"
    "fmt"
    "os/signal"
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

    pin.SetDirection(embd.Out)
    
    return pin, nil
}

func main() {
    pin, err := Init("P8_38")
    
    if err == nil {
        timeout := (time.Nanosecond/(2*uvr1611.SignalFrequency))
    	for {
            pin.PullUp()
    		time.Sleep(timeout)
            pin.PullDown()
    	}
    }
}