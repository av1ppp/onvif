package sdk

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif/errors"
)

// ReadAndParse reads the body of the http reply, unmarshals it into the "reply" object and closes
// the http reply body.
func ReadAndParse(ctx context.Context, httpReply *http.Response, reply any) error {
	data, err := readAll(ctx, httpReply)
	if err != nil {
		return err
	}

	if err := validate(httpReply, data); err != nil {
		return err
	}

	return unmarshal(data, reply)
}

// ReadAndParseWithLogging works like ReadAndParse but also logs the response body.
func ReadAndParseWithLogging(
	ctx context.Context,
	logger *logx.Logger,
	httpReply *http.Response,
	reply any,
	tag string,
) error {
	logger.Debug("reading and parsing response",
		logx.String("status", httpReply.Status),
		logx.Int("status-code", httpReply.StatusCode),
		logx.String("action", tag))

	data, err := readAll(ctx, httpReply)
	if err != nil {
		return err
	}
	logger.Debug("response was read", logx.String("response", string(data)))

	if err := validate(httpReply, data); err != nil {
		return err
	}

	return unmarshal(data, reply)
}

func unmarshal(data []byte, reply any) error {
	err := xml.Unmarshal(data, reply)
	if err != nil {
		return errors.Common.Wrap(err, "failed to decode response body")
	}
	return nil
}

func readAll(ctx context.Context, httpReply *http.Response) ([]byte, error) {
	// TODO(jfsmig): extract the deadline from ctx.Deadline() and apply it on the reply reading
	b, err := io.ReadAll(httpReply.Body)
	if err != nil {
		return nil, errors.Common.Wrap(err, "failed to read response body")
	}
	httpReply.Body.Close()
	return b, nil
}

func validate(httpReply *http.Response, data []byte) error {
	if httpReply.StatusCode < 200 || httpReply.StatusCode > 299 {
		failt := errors.FaultEnvelope{}
		if err := xml.Unmarshal(data, &failt); err != nil {
			// return errors.Common.Wrap(err, "failed to decode error response body")
			return errors.Common.New("unexpected status code (could not decode error)").
				WithProperty(errors.PropStatusCode, httpReply.StatusCode)
		}

		text := string(failt.Body.Fault.Reason.Text.Text)
		if text == "" {
			text = "<empty>"
		}

		msg := "(" + string(failt.Body.Fault.Code.Value) + ") " + text

		return errors.Soap.New(msg).
			WithProperty(errors.PropStatusCode, httpReply.StatusCode).
			WithProperty(errors.PropSoapCode, failt.Body.Fault.Code.Value)
	}
	return nil
}
