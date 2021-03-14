package main

import (
	"fmt"
	"os/exec"
)

type Brightness struct {
	CurrentValue int
	MaxValue int
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
	return exec.Command("sh", "-c", fmt.Sprintf("echo %v > %v", b.CurrentValue, CurrentDevice.GetBrightnessPath())).Run()
}

func New() (*Brightness, error) {
	max, err := GetNumberFromFile(CurrentDevice.GetMaxBrightnessPath())
	if err != nil {
		return nil, err
	}

	curr, err := GetNumberFromFile(CurrentDevice.GetBrightnessPath())
	if err != nil {
		return nil, err
	}

	return &Brightness{
		CurrentValue: curr,
		MaxValue:     max,
	}, nil
}