package main

import(
    "github.com/brutella/gouvr/uvr"
    "github.com/kidoman/embd"
    _"github.com/kidoman/embd/host/bbb"
    "os"
    "os/signal"
    _"time"
    "fmt"
    "math/big"
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
    
    if err == nil {
            
        bitReceiver := uvr.NewTestBitReceiver()
        filePath := uvr.RandomTempFilePath()
        fmt.Println("Logging to", filePath)
        logger := uvr.NewFileLogger(filePath, 5000000)
        handover := uvr.NewHandover(bitReceiver, logger)
        s := uvr.NewSignal(handover)
        
    	// clean up on exit
    	c := make(chan os.Signal, 1)
    	signal.Notify(c, os.Interrupt)
    	go func() {
    		for _ = range c {
    			fmt.Println("Flush")
                logger.Flush()
    			os.Exit(0)
    		}
    	}()
        
        err = pin.Watch(embd.EdgeBoth, func(pin embd.DigitalPin) {
            value, read_err := pin.Read()
            if read_err != nil {
                fmt.Println(read_err)
            } else {
                s.Consume(big.Word(value))
            }
        })
        if err != nil {
    	    panic(err)
        }
        
        select {
            
        }
    }
}