package main

import (
	"fmt"
	"os"

	"github.com/oxodao/brightnesshandler/brightnesshandler"
)

func main() {
	cfg, err := brightnesshandler.LoadConfig()
	if err != nil {
		panic(err)
	}

	brightness, err := brightnesshandler.New(cfg)
	if err != nil {
		panic(err)
	}

	args := brightnesshandler.ArgumentParser{Args: os.Args}
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
