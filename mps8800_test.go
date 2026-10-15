package danfysikmps8800_test

import (
	"fmt"
	"log"
	"testing"
	"time"

	danfysikmps8800 "github.com/devicehub-go/danfysik-mps8800"
	"github.com/devicehub-go/unicomm"
	"github.com/devicehub-go/unicomm/protocol/unicommtcp"
)

func TestCommunication(t *testing.T) {
	danfysik := danfysikmps8800.New(unicomm.Options{
		Protocol: unicomm.TCP,
		TCP: unicommtcp.TCPOptions{
			Host:         "10.0.28.44",
			Port:         4002,
			ReadTimeout:  200 * time.Millisecond,
			WriteTimeout: 200 * time.Millisecond,
		},
	})
	if err := danfysik.Connect(); err != nil {
		log.Fatal(err)
	}
	if err := danfysik.SelectAddress(0); err != nil {
		log.Fatal(err)
	}
	if err := danfysik.DisableAnswerMode(); err != nil {
		log.Fatal(err)
	}

	value, err := danfysik.GetFirmwareVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Firmware Version: %v\n", value)

	for channel := 0; channel <= 8; channel++ {
		value, err := danfysik.GetADC(channel)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ADC Voltage %d: %v\n", channel, value)
	}

	current, err := danfysik.GetOutputCurrent()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output Current: %v\n", current)

	mode, err := danfysik.GetControlMode()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Control Mode: %v\n", mode)

	psError, err := danfysik.GetError()
	fmt.Printf("Errors: %+v | Error: %v\n", psError, err)

	maxCurrent, err := danfysik.GetMaxRatedCurrent()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Maximum Rated Current: %v\n", maxCurrent)

	polarity, err := danfysik.GetPolarity()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Polarity: %v\n", polarity)

	demandedCurrent, err := danfysik.GetDemandedCurrent()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Demanded Current: %v\n", demandedCurrent)

	targetCurrent, err := danfysik.GetTargetEndCurrent()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Target End Current: %v\n", targetCurrent)

	rampSlope, err := danfysik.GetTargetRampSlope()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Target Ramp Slope: %v\n", rampSlope)

	status, err := danfysik.GetStatus()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Status: %+v\n", status)

	fmt.Println(danfysik.SetEndCurrent(1))
	fmt.Println(danfysik.SetRampSlope(0.3))
}
