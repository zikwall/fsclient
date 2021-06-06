package impl

import (
	"context"
	"io"
)

type FileDest struct {
	Name string
	File io.Reader
}

type Client interface {
	SendFile(context.Context, ...FileDest) error
}
