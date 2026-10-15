package protocol

import "fmt"

// Enables the answer mode for WAR and WR
func (m *MPS8800) EnableAnswerMode() error {
	if err := m.Write("ASW"); err != nil {
		return fmt.Errorf("error on enable answer mode: %v", err)
	}
	return nil
}

// Disables the answer mode for WAR and WR
func (m *MPS8800) DisableAnswerMode() error {
	if err := m.Write("NASW"); err != nil {
		return fmt.Errorf("error on disable answer mode: %v", err)
	}
	return nil
}

// Clears all active interlocks
func (m *MPS8800) ResetInterlock() error {
	if err := m.Write("RS"); err != nil {
		return fmt.Errorf("error on disable answer mode: %v", err)
	}
	return nil
}

// Switch magnet power off and resets settings of ramp and end current
func (m *MPS8800) StandBy() error {
	if err := m.Write("F"); err != nil {
		return fmt.Errorf("error on set stand by mode: %v", err)
	}
	return nil
}

// Switch off the Auxiliary line output port
func (m *MPS8800) TurnOffAuxiliaryLine() error {
	if err := m.Write("F1"); err != nil {
		return fmt.Errorf("error on turn off auxiliary line: %v", err)
	}
	return nil
}

// Turns off the addressed MPS completely, include control power
func (m *MPS8800) TurnOff() error {
	if err := m.Write("GOFF"); err != nil {
		return fmt.Errorf("error on global turn off: %v", err)
	}
	return nil
}

// Enables listen all mode where all controllers connected will
// respond to all setup and action commands
func (m *MPS8800) EnableListenAll() error {
	if err := m.Write("LALL"); err != nil {
		return fmt.Errorf("error on enable listen all mode: %v", err)
	}
	return nil
}

// Switch magnet power on
func (m *MPS8800) TurnOn() error {
	if err := m.Write("N"); err != nil {
		return fmt.Errorf("error on turn on: %v", err)
	}
	return nil
}

// Switch on the Auxiliary line output port
func (m *MPS8800) TurnOnAuxiliaryLine() error {
	if err := m.Write("N1"); err != nil {
		return fmt.Errorf("error on turn on auxiliary line: %v", err)
	}
	return nil
}

// Changes the polarity of the output voltage
func (m *MPS8800) ChangePolarity() error {
	if err := m.Write("POC"); err != nil {
		return fmt.Errorf("error on change polarity: %v", err)
	}
	return nil
}

// Stops the ramping function.
// Current remains at the value it had reached when stopped
func (m *MPS8800) Stop() error {
	if err := m.Write("STOP"); err != nil {
		return fmt.Errorf("error on stop: %v", err)
	}
	return nil
}

// Start the ramping function
func (m *MPS8800) Start() error {
	if err := m.Write("TS"); err != nil {
		return fmt.Errorf("error on stop: %v", err)
	}
	return nil
}
