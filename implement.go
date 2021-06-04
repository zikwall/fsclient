package main

import (
	"context"
	"os"
)

type Client interface {
	SendFile(context.Context, ...*os.File) error
}
