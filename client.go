package fsclient

import (
	"context"
	"fmt"
	"github.com/zikwall/fsclient/errors"
	"github.com/zikwall/fsclient/impl"
	"github.com/zikwall/fsclient/requesters"
	"net/url"
)

const (
	TypeToken = iota
	TypeBasic
	TypeJWT
)

type FsClient struct {
	uri        *url.URL
	client     impl.Client
	Uri        string
	SecureType int
	TokenType  requesters.TokenType
	Token      string
	User       string
	Password   string
}

func WithCopyFsClient(fc FsClient) (FsClient, error) {
	u, err := url.Parse(fc.Uri)

	if err != nil {
		return fc, errors.Wrap("failed parse target URL", err)
	}

	fc.uri = u

	switch fc.SecureType {
	case TypeToken:
		fc.client = requesters.NewTokenRequester(fc.uri, fc.Token, fc.TokenType)
	case TypeBasic:
		fc.client = requesters.NewBasicAuthRequester(fc.uri, fc.User, fc.Password)
	default:
		return fc, fmt.Errorf("unsupported secure type, type is: %d", fc.SecureType)
	}

	return fc, nil
}

func (fs FsClient) SendFile(context context.Context, files ...impl.FileDest) error {
	if len(files) == 0 {
		return fmt.Errorf("failed send files with empty size")
	}

	return fs.client.SendFile(context, files...)
}
