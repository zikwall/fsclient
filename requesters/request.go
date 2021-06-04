package requesters

import (
	"bytes"
	"context"
	"fmt"
	"github.com/zikwall/fsclient/errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func prepareRequest(context context.Context, uri string, files ...*os.File) (*http.Request, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	for _, file := range files {
		part, err := writer.CreateFormFile("files[]", file.Name())

		if err != nil {
			return nil, errors.Wrap("failed create form file", err)
		}

		if _, err := io.Copy(part, file); err != nil {
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
