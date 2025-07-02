package networking

import (
	"bytes"
	"context"
	"net/http"

	"github.com/av1ppp/onvif/errors"
)

// SendSoap send soap message
func SendSoap(ctx context.Context, httpClient *http.Client, endpoint, message string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(context.Background(), "POST", endpoint, bytes.NewBufferString(message))
	if err != nil {
		return nil, errors.Common.Wrap(err, "failed to create POST request")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return resp, errors.Common.Wrap(err, "failed to send POST request")
	}

	return resp, nil
}
