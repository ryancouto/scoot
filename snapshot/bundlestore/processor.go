package bundlestore

import (
	"io"
)

type Processor interface {
	Process(input io.ReadCloser) (io.ReadCloser, error)
}

type IdentityProcessor struct{}

func (p *IdentityProcessor) Process(input io.ReadCloser) (io.ReadCloser, error) {
	return input, nil
}
