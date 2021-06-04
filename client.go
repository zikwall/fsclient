package main

import (
	"context"
	"fmt"
	"github.com/zikwall/fsclient/errors"
	"github.com/zikwall/fsclient/requesters"
	"log"
	"net/url"
	"os"
)

const (
	TypeToken = iota
	TypeBasic
	TypeJWT
)

type FsClient struct {
	uri        *url.URL
	client     Client
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

func (fs FsClient) SendFile(context context.Context, files ...*os.File) error {
	if len(files) == 0 {
		return fmt.Errorf("failed send files with empty size")
	}

	return fs.client.SendFile(context, files...)
}

func main() {
	fsclient, err := WithCopyFsClient(FsClient{
		Uri:        "http://localhost:1337/",
		SecureType: TypeToken,
		TokenType:  requesters.TokenTypeQuery,
		Token:      "changemeplease123",
		User:       "qwx1337",
		Password:   "123456",
	})

	if err != nil {
		log.Fatal(err)
	}

	f1, _ := os.Open("test.txt")
	f2, _ := os.Open("test2.txt")

	defer func() {
		f1.Close()
		f2.Close()
	}()

	if err := fsclient.SendFile(context.Background(), f1, f2); err != nil {
		log.Fatal(err)
	}
}
