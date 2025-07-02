// Code generated : DO NOT EDIT.

package analytics

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/analytics"
	"github.com/av1ppp/onvif/errors"
)

// GetSupportedMetadata forwards the call to onvif.Do then parses the payload of the reply as a GetSupportedMetadataResponse.
func GetSupportedMetadata(ctx context.Context, dev *onvif.Device, request *onvif.Req[analytics.GetSupportedMetadata]) (analytics.GetSupportedMetadataResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSupportedMetadataResponse analytics.GetSupportedMetadataResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetSupportedMetadataResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSupportedMetadata")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetSupportedMetadata")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetSupportedMetadataResponse, err
}
