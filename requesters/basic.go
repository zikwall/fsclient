package requesters

import (
	"context"
	"github.com/zikwall/fsclient/impl"
	"net/url"
)

type BasicAuthRequester struct {
	uri      string
	user     string
	password string
}

func NewBasicAuthRequester(uri *url.URL, user, password string) BasicAuthRequester {
	return BasicAuthRequester{
		uri:      uri.String(),
		user:     user,
		password: password,
	}
}

func (br BasicAuthRequester) SendFile(context context.Context, files ...impl.FileDest) error {
	request, err := prepareRequest(context, br.uri, files...)

	if err != nil {
		return err
	}

	request.SetBasicAuth(br.user, br.password)

	return prepareResponse(context, request)
}
