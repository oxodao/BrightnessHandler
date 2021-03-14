package main

import (
	"fmt"
	"os"
)

var CurrentDevice Device

func main() {
	CurrentDevice = NewDevice("amdgpu_bl0", "brightness", "max_brightness")

	brightness, err := New()
	if err != nil {
		panic(err)
	}

	err = brightness.EnforcePermissions()
	if err != nil {
		panic(err)
	}

	args := ArgumentParser{Args: os.Args}
	err = args.Parse(brightness)
	if err != nil {
		help()
		return
	}

	switch args.Command {
	case "get":
	case "add":
		brightness.Add(args.WantedBrightness)
	case "sub":
		brightness.Sub(args.WantedBrightness)
	case "set":
		brightness.Set(args.WantedBrightness)
	default:
		help()
	}

	if args.Command == "add" || args.Command == "sub" || args.Command == "set" {
		brightness.Write()
	}

	fmt.Printf("%v%%\n", brightness.GetCurrent())
}

func help() {
	fmt.Println("BrightnessHandler usage")
	fmt.Println("-----------------------")
	fmt.Println("\t- get")
	fmt.Println("\t- set PERCENTAGE")
	fmt.Println("\t- add PERCENTAGE")
	fmt.Println("\t- sub PERCENTAGE")
}