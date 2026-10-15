package danfysikmps8800

import (
	"github.com/devicehub-go/danfysik-mps8800/protocol"
	"github.com/devicehub-go/unicomm"
)

// Creates a new instance of Danfysik MPS 8800 Power Supply
func New(options unicomm.Options) *protocol.MPS8800 {
	options.Delimiter = "\r"
	return &protocol.MPS8800{
		Communication: unicomm.New(options),
	}
}
