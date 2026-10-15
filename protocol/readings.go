package protocol

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MPSError struct {
	Code        string
	Description string
}

type MPSInterlock struct {
	MagnetFault                   bool `json:"magnet-fault"`
	ExternalSpare1                bool `json:"external-spare-01"`
	ExternalSpare2                bool `json:"external-spare-02"`
	FreewheelDiodeOverTemperature bool `json:"freeewheel-diode-overtemp"`
	LowWaterFlow                  bool `json:"low-water-flow"`
	DoorOpen                      bool `json:"door-open"`
	DiodeHeatSinkOverTemperature  bool `json:"diode-heatsink-overtemp"`
	ChassisOverTemperature        bool `json:"chassis-overtemp"`
	IGBTHeatSinkOverTemperature   bool `json:"igbt-heatsink-overtemp"`
	HFDiodeOverTemperature        bool `json:"hf-diode-overtemp"`
	SwitchRegulatorDCCTFailure    bool `json:"switchreg-dcct-failure"`
	SwitchRegulatorSupplyFailure  bool `json:"switchreg-supply-failure"`
	IGBTDriverFailure             bool `json:"igbt-driver-failure"`
	OverCurrent                   bool `json:"overcurrent"`
}

type MPSWarning struct {
	ACUnderVoltage         bool `json:"ac-undervolt"`
	ExcessiveCurrentRipple bool `json:"excessive-current-ripple"`
	GroundLeakCurrent      bool `json:"ground-leak-current"`
}

type MPSStatus struct {
	PositivePolarity          bool `json:"positive-polarity"`
	NegativePolarity          bool `json:"negative-polarity"`
	ExternalInternalReference bool `json:"ext/int-referece"`
	RemoteControlIndication   bool `json:"remote-control"`
	MainPowerIsOn             bool `json:"main-power-on"`
	Ready                     bool `json:"ready"`
}

type StatusBits struct {
	Status    MPSStatus    `json:"status"`
	Warning   MPSWarning   `json:"warning"`
	Interlock MPSInterlock `json:"interlock"`
}

func (s StatusBits) ToJson() ([]byte, error) {
	return json.Marshal(s)
}

// Gets the firmware version
func (m *MPS8800) GetFirmwareVersion() (string, error) {
	value, err := m.Query("VER")
	if err != nil {
		return "", fmt.Errorf("error on get firmare version: %v", err)
	}
	return value, nil
}

// Gets the analog measurement from an AD converter
func (m *MPS8800) GetADC(channel int) (float64, error) {
	if 0 > channel || channel > 10 {
		return 0, fmt.Errorf("ADC channel must be between 0-8, got %d", channel)
	}
	value, err := m.QueryFloat64(fmt.Sprintf("AD %d", channel))
	if err != nil {
		return 0, fmt.Errorf("error on get ADC %d: %v", channel, err)
	}
	return value, nil
}

// Gets the R-AC voltage in V
func (m *MPS8800) GetVoltageR() (float64, error) {
	return m.GetADC(0)
}

// Gets the S-AC voltage in V
func (m *MPS8800) GetVoltageS() (float64, error) {
	return m.GetADC(1)
}

// Gets the T-AC voltage in V
func (m *MPS8800) GetVoltageT() (float64, error) {
	return m.GetADC(2)
}

// Gets the R-AC current in A
func (m *MPS8800) GetCurrentR() (float64, error) {
	value, err := m.GetADC(3)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// Gets the S-AC current in A
func (m *MPS8800) GetCurrentS() (float64, error) {
	value, err := m.GetADC(4)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// Gets the T-AC current in A
func (m *MPS8800) GetCurrentT() (float64, error) {
	value, err := m.GetADC(5)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// Gets the output voltage in V
func (m *MPS8800) GetOutputVoltage() (float64, error) {
	return m.GetADC(6)
}

// Gets the ground leak current in A
func (m *MPS8800) GetGroundLeakCurrent() (float64, error) {
	value, err := m.GetADC(7)
	if err != nil {
		return 0, err
	}
	return value / 1000.0, nil
}

// Gets the delta temperature in celsius degree
func (m *MPS8800) GetDeltaTemperature() (float64, error) {
	return m.GetADC(8)
}

// Gets the actual output current in A
func (m *MPS8800) GetOutputCurrent() (float64, error) {
	value, err := m.QueryFloat64("ADCV")
	if err != nil {
		return 0, fmt.Errorf("error on get output current: %v", err)
	}
	return value / 1000.0, nil
}

// Gets the control mode (local/remote)
func (m *MPS8800) GetControlMode() (string, error) {
	value, err := m.Query("CMDSTATE")
	if err != nil {
		return "", fmt.Errorf("error on get control mode: %v", err)
	}
	return value, nil
}

// Gets the error code (not works yet)
func (m *MPS8800) GetError() (MPSError, error) {
	errorCode, err := m.Query("ERRC")
	if strings.Contains(err.Error(), "timeout") {
		return MPSError{}, nil
	}
	if err != nil {
		return MPSError{}, fmt.Errorf("error on get error code: %v", err)
	}
	errorDescription, err := m.Query("ERRT")
	if err != nil {
		return MPSError{}, fmt.Errorf("error on get error description: %v", err)
	}
	return MPSError{
		Code:        errorCode,
		Description: errorDescription,
	}, nil
}

// Gets the maximum rated current in A
func (m *MPS8800) GetMaxRatedCurrent() (float64, error) {
	value, err := m.QueryFloat64("MAX")
	if err != nil {
		return 0, fmt.Errorf("error on get maximum rated current: %v", err)
	}
	return value / 1000.0, nil
}

// Gets the power supply polarity
func (m *MPS8800) GetPolarity() (string, error) {
	value, err := m.Query("PO")
	if err != nil {
		return "", fmt.Errorf("error on get control mode: %v", err)
	}
	return value, nil
}

// Gets the actual demanded output current in A
func (m *MPS8800) GetDemandedCurrent() (float64, error) {
	value, err := m.QueryFloat64("RA")
	if err != nil {
		return 0, fmt.Errorf("error on get demanded current: %v", err)
	}
	return value / 1000.0, nil
}

// Gets the configured end-current setpoint in A
func (m *MPS8800) GetTargetEndCurrent() (float64, error) {
	value, err := m.QueryFloat64("RAR")
	if err != nil {
		return 0, fmt.Errorf("error on get target end current: %v", err)
	}
	return value / 1000.0, nil
}

// Gets the configured ramp slope in A/s
func (m *MPS8800) GetTargetRampSlope() (float64, error) {
	value, err := m.QueryFloat64("RR")
	if err != nil {
		return 0, fmt.Errorf("error on get target end current: %v", err)
	}
	maxRated, err := m.GetMaxRatedCurrent()
	if err != nil {
		return 0, fmt.Errorf("error on get target end current: %v", err)
	}
	return value * maxRated * 0.001, nil
}

// Gets the ramp

// Gets the internal device status
func (m *MPS8800) GetStatus() (StatusBits, error) {
	value, err := m.Query("S1")
	if err != nil {
		return StatusBits{}, fmt.Errorf("error on get internal status: %v", err)
	}
	if len(value) != 32 {
		return StatusBits{}, fmt.Errorf(
			"invalid status length: expected 32, got %d", len(value),
		)
	}
	return StatusBits{
		Status: MPSStatus{
			PositivePolarity:          value[10] == '!',
			NegativePolarity:          value[11] == '!',
			ExternalInternalReference: value[24] == '!',
			RemoteControlIndication:   value[26] == '!',
			MainPowerIsOn:             value[30] == '!',
			Ready:                     value[31] == '!',
		},
		Warning: MPSWarning{
			ACUnderVoltage:         value[25] == '!',
			ExcessiveCurrentRipple: value[27] == '!',
			GroundLeakCurrent:      value[28] == '!',
		},
		Interlock: MPSInterlock{
			MagnetFault:                   value[1] == '!',
			ExternalSpare1:                value[2] == '!',
			ExternalSpare2:                value[3] == '!',
			FreewheelDiodeOverTemperature: value[7] == '!',
			LowWaterFlow:                  value[8] == '!',
			DoorOpen:                      value[9] == '!',
			DiodeHeatSinkOverTemperature:  value[16] == '!',
			ChassisOverTemperature:        value[17] == '!',
			IGBTHeatSinkOverTemperature:   value[18] == '!',
			HFDiodeOverTemperature:        value[19] == '!',
			SwitchRegulatorDCCTFailure:    value[20] == '!',
			SwitchRegulatorSupplyFailure:  value[21] == '!',
			IGBTDriverFailure:             value[22] == '!',
			OverCurrent:                   value[29] == '!',
		},
	}, nil
}
