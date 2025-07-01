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

// Call_GetMetadataConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetMetadataConfigurationResponse.
func Call_GetMetadataConfiguration(ctx context.Context, dev *onvif.Device, request media.GetMetadataConfiguration) (media.GetMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetadataConfigurationResponse media.GetMetadataConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetMetadataConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetMetadataConfiguration")
	}
}

// CallWithLogging_GetMetadataConfiguration works like Call_GetMetadataConfiguration but also logs the response body.
func CallWithLogging_GetMetadataConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetMetadataConfiguration) (media.GetMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetMetadataConfigurationResponse media.GetMetadataConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetMetadataConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetMetadataConfiguration")
		return reply.Body.GetMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetMetadataConfiguration")
	}
}
