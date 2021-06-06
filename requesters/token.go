package requesters

import (
	"context"
	"fmt"
	"github.com/zikwall/fsclient/impl"
	"net/url"
)

type TokenType int

const (
	TokenTypeHeader TokenType = iota
	TokenTypeQuery
)

type TokenRequester struct {
	token     string
	tokenType TokenType
	uri       string
}

func NewTokenRequester(uri *url.URL, token string, tokenType ...TokenType) TokenRequester {
	tr := TokenRequester{}
	tr.token = token
	tr.tokenType = TokenTypeHeader

	if len(tokenType) > 0 {
		tr.tokenType = tokenType[0]
	}

	if tr.tokenType == TokenTypeQuery {
		q := uri.Query()
		q.Set("token", tr.token)
		uri.RawQuery = q.Encode()
	}

	tr.uri = uri.String()

	return tr
}

func (tr TokenRequester) SendFile(context context.Context, files ...impl.FileDest) error {
	request, err := prepareRequest(context, tr.uri, files...)

	if err != nil {
		return err
	}

	if tr.tokenType == TokenTypeHeader {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tr.token))
	}

	return prepareResponse(context, request)
}
