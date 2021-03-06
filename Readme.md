# GOUVR

This project provides a boilerplate to read data from the data bus from devices manufactured by TA (Technische Alternative)

- UVR 31  
- UVR 42  
- UVR 64  
- HRZ 65  
- EEG 30  
- TFM 66  
- [UVR 1611][UVR1611-Website] available under `uvr/1611`
- UVR 61-3
- ESR 21

The data bus uses a signal at 488 frequency to transmit bits. The bit stream is manchester encoded as described in `docs/Schnittstelle Datenleitung 1.6.pdf` (German only).

## UVR1611

The project contains an implementation of a pipeline to decode the data signal of an UVR1611. The pipeline consists of several stages which decode bits to bytes and then to packets. The stages are abstracted by interfaces

    WordConsumer -> BitConsumer -> ByteConsumer -> PacketConsumer


### Pipeline Stages

1. `WordConsumer` in `signal.go`: Converts a bit of type `big.Word` to `Bit` which includes the bit value and a timestamp
2. `BitConsumer` in `uvr/1611/sync.go`: Synchronizes the decoding and passes bits through to next bit consumer when synchronized
3. `BitConsumer` in `byte.go`: Accumulates bits to one byte and checks start and stop bit
4. `ByteConsumer` in `uvr/1611/packet_decoder.go`: Accumulates bytes to one packet
5. `PacketConsumer` by `packetConsumer` in `uvr/1611/util.go` which calls a function when packet arrives

## TODO

- Speed step `speed_step.go`
	- Add methods to check if enabled
	- Add method to get value
- Heat meter `heatmeter.go`
	- Decode power of heatmeters

[UVR1611-Website]: http://www.ta.co.at/en/products/uvr1611/

# Contact

Matthias Hochgatterer

Github: [https://github.com/brutella/](https://github.com/brutella/)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)


# License

gouvr is available under the MIT license. See the LICENSE file for more info.