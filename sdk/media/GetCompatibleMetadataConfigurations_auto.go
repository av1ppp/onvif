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

// Call_GetCompatibleMetadataConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleMetadataConfigurationsResponse.
func Call_GetCompatibleMetadataConfigurations(ctx context.Context, dev *onvif.Device, request media.GetCompatibleMetadataConfigurations) (media.GetCompatibleMetadataConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleMetadataConfigurationsResponse media.GetCompatibleMetadataConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleMetadataConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleMetadataConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetCompatibleMetadataConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleMetadataConfigurations")
	}
}

// CallWithLogging_GetCompatibleMetadataConfigurations works like Call_GetCompatibleMetadataConfigurations but also logs the response body.
func CallWithLogging_GetCompatibleMetadataConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetCompatibleMetadataConfigurations) (media.GetCompatibleMetadataConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleMetadataConfigurationsResponse media.GetCompatibleMetadataConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleMetadataConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleMetadataConfigurations")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCompatibleMetadataConfigurations")
		return reply.Body.GetCompatibleMetadataConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleMetadataConfigurations")
	}
}
