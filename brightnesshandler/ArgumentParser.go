package brightnesshandler

import (
	"errors"
	"strconv"
)

type ArgumentParser struct {
	Args             []string
	Command          string
	WantedBrightness int
}

func (sc *ArgumentParser) Parse(b *Brightness) error {
	if len(sc.Args) != 2 && len(sc.Args) != 3 {
		return errors.New("not enough args")
	}

	if !isAllowed(sc.Args[1], "get", "set", "add", "sub") {
		return errors.New("Not allowed command: " + sc.Args[1])
	}

	sc.Command = sc.Args[1]

	if len(sc.Args) == 3 {
		wantedValue, err := strconv.Atoi(sc.Args[2])
		if err != nil {
			return err
		}

		sc.WantedBrightness = int((float64(wantedValue) / 100) * float64(b.MaxValue))
	}

	return nil
}

func isAllowed(arg string, allowedValues ...string) bool {
	for _, s := range allowedValues {
		if arg == s {
			return true
		}
	}

	return false
}
