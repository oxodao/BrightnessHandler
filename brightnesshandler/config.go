package brightnesshandler

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Device string `json:"device"`
}

func LoadConfig() (defaultConfig *Config, err error) {
	defaultConfig = &Config{
		Device: "amdgpu_bl0",
	}

	device := os.Getenv("BH_DEVICE")
	if len(device) > 0 {
		defaultConfig.Device = device
		return
	}

	// Probably useless as the user needs to be logged as root
	// and no one in their right minds will have some userland config in the root user
	// People should just use `sudo BH_DEVICE=amdgpu_bl0 bh set 100`
	data, err := ioutil.ReadFile(getPath() + "/config.json")
	if err != nil {
		return defaultConfig, nil
	}

	err = json.Unmarshal(data, defaultConfig)

	return
}

func (c *Config) GetBrightnessPath() string {
	return "/sys/class/backlight/" + c.Device + "/brightness"
}

func (c *Config) GetMaxBrightnessPath() string {
	return "/sys/class/backlight/" + c.Device + "/max_brightness"
}

func getPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "./"
	}

	configPath := dirname + "/.config/brightnesshandler"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return "./"
	}

	return configPath + "/"
}
