// This app writes the output of a GPIO pin to the console or file
package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/signal"

	"github.com/brutella/gouvr/uvr"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
)

func Init(file string) (embd.DigitalPin, error) {
	embd.InitGPIO()
	pin, pin_err := embd.NewDigitalPin(file)
	if pin_err != nil {
		fmt.Printf("Error opening pin! %s\n", pin_err)
		return nil, pin_err
	}

	pin.SetDirection(embd.In)

	return pin, nil
}

func main() {
	var (
		to   = flag.String("to", "console", "Log to console or file")
		file = flag.String("file", "", "Log file")
		port = flag.String("port", "P8_07", "GPIO port; default P8_07")
	)

	flag.Parse()

	pin, err := Init(*port)

	if err == nil {

		var logger Logger

		switch *to {
		case "console":
			log.Println("Logging to console")
			logger = uvr.NewConsoleLogger()
		case "file":
			filePath = uvr.RandomTempFilePath()
			if len(file) > 0 {
				filePath = file
			}
			log.Println("Logging to", filePath)
			logger = NewFileLogger(filePath, 10000)
		default:
			log.Fatal("Invalid to flag", *to)
		}

		bitReceiver := uvr.NewTestBitReceiver()
		handover := uvr.NewHandover(bitReceiver, logger)
		s := uvr.NewSignal(handover)

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

		// clean up on exit
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for _ = range c {
				fmt.Println("Flush")
				logger.Flush()
				pin.Close()
				embd.CloseGPIO()
				os.Exit(0)
			}
		}()

		select {}
	}
}
