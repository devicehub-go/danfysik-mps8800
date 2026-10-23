package protocol

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/devicehub-go/unicomm"
)

type MPS8800 struct {
	Address       int
	Communication unicomm.Unicomm
	mutex         sync.Mutex
}

// Establishes a connection with power supply
func (m *MPS8800) Connect() error {
	if err := m.Communication.Connect(); err != nil {
		return fmt.Errorf("error on connect: %v", err)
	}
	return nil
}

// Closes the connection with power supply
func (m *MPS8800) Disconnect() error {
	if err := m.Communication.Disconnect(); err != nil {
		return fmt.Errorf("error on close connection with MPS8800: %v", err)
	}
	return nil
}

// Checks if there is a connection established with power supply
func (m *MPS8800) IsConnected() bool {
	return m.Communication.IsConnected()
}

func (m *MPS8800) writeUnsafe(message string) error {
	if !m.Communication.IsConnected() {
		return fmt.Errorf("no connection established with MPS8800")
	}
	if !strings.HasSuffix(message, "\r") {
		message = message + "\r"
	}
	if err := m.Communication.Write([]byte(message)); err != nil {
		return fmt.Errorf("error on query to MPS8800: %v", err)
	}
	return nil
}

// Writes a message to power supply
func (m *MPS8800) Write(message string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.writeUnsafe(message)
}

// Queries a message from power supply
func (m *MPS8800) Query(message string) (string, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if err := m.writeUnsafe(message); err != nil {
		return "", fmt.Errorf("error on query to MPS8800: %v", err)
	}
	response, err := m.Communication.Read(1024)
	if err != nil {
		return "", fmt.Errorf("error on query to MPS8800: %v", err)
	}

	responseStr := string(response)
	responseStr = strings.TrimSuffix(responseStr, "\r")
	responsesArr := strings.Split(responseStr, "\r")
	// This allows handle with possible asynchronous error messages
	if len(responsesArr) > 1 {
		responseStr = responsesArr[len(responsesArr)-1]
	}
	if len(responsesArr) > 2 {
		log.Printf("[DEBUG] Some messages were ignored: %v", responsesArr)
	}

	responseStr = strings.ReplaceAll(responseStr, "\r", "")
	return responseStr, nil
}

// Queries a float64 number from the controller
func (m *MPS8800) QueryFloat64(command string) (float64, error) {
	response, err := m.Query(command)
	if err != nil {
		return 0, err
	}
	value, err := strconv.ParseFloat(response, 64)
	if err != nil {
		return 0, fmt.Errorf("error on parse query response '%s' to float64: %v", response, err)
	}
	return value, nil
}

// Queries an int number from the controller
func (m *MPS8800) QueryInt(command string) (int, error) {
	response, err := m.Query(command)
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(response)
	if err != nil {
		return 0, fmt.Errorf("error on parse query response '%s' to int: %v", response, err)
	}
	return value, nil
}
