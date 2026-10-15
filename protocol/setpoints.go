package protocol

import (
	"fmt"
	"math"
)

// Select power supply address
func (m *MPS8800) SelectAddress(address int) error {
	if 0 > address || address > 255 {
		return fmt.Errorf("address must be between 0-255, got %d", address)
	}
	if err := m.Write(fmt.Sprintf("ADR %03d", address)); err != nil {
		return fmt.Errorf("error on set address: %v", err)
	}
	return nil
}

// Sets the end current in A
func (m *MPS8800) SetEndCurrent(endCurrent float64) error {
	maxRated, err := m.GetMaxRatedCurrent()
	if err != nil {
		return fmt.Errorf("error on set end current: %v", err)
	}

	mAcurrent := math.Round(endCurrent * 1000)
	if 0 > endCurrent || endCurrent > maxRated {
		return fmt.Errorf(
			"end current must be between 0-%.3f, got %.3f",
			maxRated, endCurrent,
		)
	}
	if err := m.Write(fmt.Sprintf("WAR %06.0f", mAcurrent)); err != nil {
		return fmt.Errorf("error on set end current: %v", err)
	}

	actual, err := m.GetTargetEndCurrent()
	if err != nil {
		return fmt.Errorf("error on set end current: %v", err)
	}
	if actual != endCurrent {
		return fmt.Errorf("error on set end current: target and actual mismatch")
	}
	return nil
}

// Sets the slope rate for the ramping function in A/s
func (m *MPS8800) SetRampSlope(rate float64) error {
	maxRated, err := m.GetMaxRatedCurrent()
	if err != nil {
		return fmt.Errorf("error on set end current: %v", err)
	}

	percentRate := rate / (maxRated * 0.001)
	if 1 > percentRate || percentRate > 100 {
		return fmt.Errorf(
			"ramp slope must be between %.1f-%.1f, got %f",
			maxRated*0.01, maxRated, rate,
		)
	}
	if err := m.Write(fmt.Sprintf("WR %03.0f", percentRate)); err != nil {
		return fmt.Errorf("error on set ramp rate: %v", err)
	}

	actual, err := m.GetTargetRampSlope()
	if err != nil {
		return fmt.Errorf("error on set ramp rate: %v", err)
	}
	if actual != rate {
		return fmt.Errorf("error on set ramp rate: target and actual mismatch")
	}
	return nil
}
