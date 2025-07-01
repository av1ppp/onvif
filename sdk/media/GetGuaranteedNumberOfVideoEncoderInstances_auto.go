// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetGuaranteedNumberOfVideoEncoderInstances forwards the call to dev.CallMethod() then parses the payload of the reply as a GetGuaranteedNumberOfVideoEncoderInstancesResponse.
func Call_GetGuaranteedNumberOfVideoEncoderInstances(ctx context.Context, dev *onvif.Device, request media.GetGuaranteedNumberOfVideoEncoderInstances) (media.GetGuaranteedNumberOfVideoEncoderInstancesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetGuaranteedNumberOfVideoEncoderInstancesResponse media.GetGuaranteedNumberOfVideoEncoderInstancesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetGuaranteedNumberOfVideoEncoderInstancesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetGuaranteedNumberOfVideoEncoderInstances")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetGuaranteedNumberOfVideoEncoderInstancesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetGuaranteedNumberOfVideoEncoderInstances")
	}
}

// CallWithLogging_GetGuaranteedNumberOfVideoEncoderInstances works like Call_GetGuaranteedNumberOfVideoEncoderInstances but also logs the response body.
func CallWithLogging_GetGuaranteedNumberOfVideoEncoderInstances(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetGuaranteedNumberOfVideoEncoderInstances) (media.GetGuaranteedNumberOfVideoEncoderInstancesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetGuaranteedNumberOfVideoEncoderInstancesResponse media.GetGuaranteedNumberOfVideoEncoderInstancesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetGuaranteedNumberOfVideoEncoderInstancesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetGuaranteedNumberOfVideoEncoderInstances")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetGuaranteedNumberOfVideoEncoderInstances")
		return reply.Body.GetGuaranteedNumberOfVideoEncoderInstancesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetGuaranteedNumberOfVideoEncoderInstances")
	}
}
