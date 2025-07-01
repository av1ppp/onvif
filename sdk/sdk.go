package sdk

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"

	"github.com/av1ppp/logx"
	"github.com/av1ppp/onvif/errors"
)

func ReadAndParse(ctx context.Context, httpReply *http.Response, reply interface{}, tag string) error {
	// TODO(jfsmig): extract the deadline from ctx.Deadline() and apply it on the reply reading
	b, err := io.ReadAll(httpReply.Body)
	if err != nil {
		return errors.Common.Wrap(err, "failed to read response body")
	}

	httpReply.Body.Close()

	err = xml.Unmarshal(b, reply)
	if err != nil {
		return errors.Common.Wrap(err, "failed to decode response body")
	}
	return nil
}

func ReadAndParseWithLogging(ctx context.Context, logger *logx.Logger, httpReply *http.Response, reply interface{}, tag string) error {
	logger.Debug("reading and parsing response",
		logx.String("status", httpReply.Status),
		logx.Int("status-code", httpReply.StatusCode),
		logx.String("action", tag))

	return ReadAndParse(ctx, httpReply, reply, tag)
}
