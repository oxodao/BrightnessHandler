package brightnesshandler

import (
	"fmt"
	"os"
)

type Brightness struct {
	Config       *Config
	CurrentValue int
	MaxValue     int
}

func (b *Brightness) Add(wanted int) {
	wanted = b.CurrentValue + wanted

	if wanted > b.MaxValue {
		wanted = b.MaxValue
	}

	b.CurrentValue = wanted
}

func (b *Brightness) Sub(wanted int) {
	wanted = b.CurrentValue - wanted

	if wanted < 0 {
		wanted = 0
	}

	b.CurrentValue = wanted
}

func (b *Brightness) Set(wanted int) {
	if wanted < 0 {
		wanted = 0
	}

	if wanted > b.MaxValue {
		wanted = b.MaxValue
	}

	b.CurrentValue = wanted
}

func (b *Brightness) GetCurrent() int {
	return int((float64(b.CurrentValue) / float64(b.MaxValue)) * 100)
}

func (b *Brightness) Write() error {
	return os.WriteFile(b.Config.GetBrightnessPath(), []byte(fmt.Sprintf("%v", b.CurrentValue)), 0700)
}

func New(cfg *Config) (*Brightness, error) {
	max, err := GetNumberFromFile(cfg.GetMaxBrightnessPath())
	if err != nil {
		return nil, err
	}

	curr, err := GetNumberFromFile(cfg.GetBrightnessPath())
	if err != nil {
		return nil, err
	}

	return &Brightness{
		Config:       cfg,
		CurrentValue: curr,
		MaxValue:     max,
	}, nil
}
