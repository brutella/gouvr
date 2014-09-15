package uvr

import(
    "fmt"
)

type speedStep struct {
    value Byte
}

func NewSpeedStep(value Byte) speedStep {
    return speedStep{value: value}
}

func (s speedStep) ToString() string {
    var disabled bool = s.value & 0x80 == 0x80 // 1000 0000
    if disabled == true {
        return "?"
    }
    
    var value = s.value & 0x1F            // 0001 1111
    return fmt.Sprintf("%d", value)
}