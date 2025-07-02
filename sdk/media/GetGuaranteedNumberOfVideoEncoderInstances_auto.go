// Code generated : DO NOT EDIT.

package sdkmedia

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetGuaranteedNumberOfVideoEncoderInstances forwards the call to onvif.Do then parses the payload of the reply as a GetGuaranteedNumberOfVideoEncoderInstancesResponse.
func GetGuaranteedNumberOfVideoEncoderInstances(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetGuaranteedNumberOfVideoEncoderInstances]) (media.GetGuaranteedNumberOfVideoEncoderInstancesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetGuaranteedNumberOfVideoEncoderInstancesResponse media.GetGuaranteedNumberOfVideoEncoderInstancesResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetGuaranteedNumberOfVideoEncoderInstancesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetGuaranteedNumberOfVideoEncoderInstances")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetGuaranteedNumberOfVideoEncoderInstances")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetGuaranteedNumberOfVideoEncoderInstancesResponse, err
}
