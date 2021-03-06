package uvr

type DeviceType Byte

const (
	DeviceTypeUVR31   = DeviceType(0x30)
	DeviceTypeUVR42   = DeviceType(0x10)
	DeviceTypeUVR64   = DeviceType(0x20)
	DeviceTypeHRZ65   = DeviceType(0x60)
	DeviceTypeEEG30   = DeviceType(0x50)
	DeviceTypeTFM66   = DeviceType(0x40)
	DeviceTypeUVR1611 = DeviceType(0x80)
	DeviceTypeUVR61_3 = DeviceType(0x90)
	DeviceTypeESR21   = DeviceType(0x70)
)

// DeviceTypeToString returns a string representation of the device type
func DeviceTypeToString(deviceType DeviceType) string {
	switch deviceType {
	case DeviceTypeUVR31:
		return "UVR 31"
	case DeviceTypeUVR42:
		return "UVR 42"
	case DeviceTypeUVR64:
		return "UVR 64"
	case DeviceTypeHRZ65:
		return "HRZ 65"
	case DeviceTypeEEG30:
		return "EEG 30"
	case DeviceTypeTFM66:
		return "TFM 66"
	case DeviceTypeUVR1611:
		return "UVR 1611"
	case DeviceTypeUVR61_3:
		return "UVR 61-3"
	case DeviceTypeESR21:
		return "ESR 21"
	}

	return "Unknown"
}
