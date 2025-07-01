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

// Call_GetMetadataConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetMetadataConfigurationOptionsResponse.
func Call_GetMetadataConfigurationOptions(ctx context.Context, dev *onvif.Device, request media.GetMetadataConfigurationOptions) (media.GetMetadataConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetadataConfigurationOptionsResponse media.GetMetadataConfigurationOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetMetadataConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetMetadataConfigurationOptions")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetMetadataConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetMetadataConfigurationOptions")
	}
	return reply.Body.GetMetadataConfigurationOptionsResponse, nil
}

// CallWithLogging_GetMetadataConfigurationOptions works like Call_GetMetadataConfigurationOptions but also logs the response body.
func CallWithLogging_GetMetadataConfigurationOptions(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetMetadataConfigurationOptions) (media.GetMetadataConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetadataConfigurationOptionsResponse media.GetMetadataConfigurationOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetMetadataConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetMetadataConfigurationOptions")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetMetadataConfigurationOptions")
	if err != nil {
		return reply.Body.GetMetadataConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetMetadataConfigurationOptions")
	}
	return reply.Body.GetMetadataConfigurationOptionsResponse, nil
}
