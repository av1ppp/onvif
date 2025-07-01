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

// Call_GetMetadataConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetMetadataConfigurationsResponse.
func Call_GetMetadataConfigurations(ctx context.Context, dev *onvif.Device, request media.GetMetadataConfigurations) (media.GetMetadataConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetadataConfigurationsResponse media.GetMetadataConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetMetadataConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetMetadataConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetMetadataConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetMetadataConfigurations")
	}
}

// CallWithLogging_GetMetadataConfigurations works like Call_GetMetadataConfigurations but also logs the response body.
func CallWithLogging_GetMetadataConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetMetadataConfigurations) (media.GetMetadataConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetadataConfigurationsResponse media.GetMetadataConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetMetadataConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetMetadataConfigurations")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetMetadataConfigurations")
		return reply.Body.GetMetadataConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetMetadataConfigurations")
	}
}
