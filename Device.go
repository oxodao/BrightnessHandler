package main

type Device struct {
	name              string
	currentBrightness string
	maxBrightness     string
}

func NewDevice(name, currentBrightness, maxBrightness string) Device {
	return Device{
		name:              name,
		currentBrightness: currentBrightness,
		maxBrightness:     maxBrightness,
	}
}

func (d *Device) GetBrightnessPath() string {
	return "/sys/class/backlight/" + d.name + "/" + d.currentBrightness
}

func (d *Device) GetMaxBrightnessPath() string {
	return "/sys/class/backlight/" + d.name + "/" + d.maxBrightness
}