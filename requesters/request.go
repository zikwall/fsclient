package requesters

import (
	"bytes"
	"context"
	"fmt"
	"github.com/zikwall/fsclient/errors"
	"github.com/zikwall/fsclient/impl"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func prepareRequest(context context.Context, uri string, dests ...impl.FileDest) (*http.Request, error) {
	var body bytes.Buffer
	var err error

	writer := multipart.NewWriter(&body)

	for _, dest := range dests {
		filename := dest.Name
		switch f := dest.File.(type) {
		case *os.File:
			if filename == "" {
				filename = f.Name()
			}
		}

		part, err := writer.CreateFormFile("files[]", filename)

		if err != nil {
			return nil, errors.Wrap("failed create form file", err)
		}

		if _, err := io.Copy(part, dest.File); err != nil {
			return nil, errors.Wrap("failed copy buffer", err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, errors.Wrap("failed close buffer writter", err)
	}

	request, err := http.NewRequestWithContext(context, http.MethodPost, uri, &body)

	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Content-Length", fmt.Sprintf("%d", body.Len()))

	return request, nil
}

func prepareResponse(context context.Context, request *http.Request) error {
	// https://github.com/golang/go/issues/36095
	clone := request.Clone(context)
	clone.Body, _ = request.GetBody()

	response, err := (&http.Client{}).Do(clone)

	if err != nil {
		return errors.Wrap("failed prepare response", err)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid http status, code is: %d", response.StatusCode)
	}

	return nil
}
