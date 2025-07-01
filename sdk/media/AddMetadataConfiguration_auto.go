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

// Call_AddMetadataConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddMetadataConfigurationResponse.
func Call_AddMetadataConfiguration(ctx context.Context, dev *onvif.Device, request media.AddMetadataConfiguration) (media.AddMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddMetadataConfigurationResponse media.AddMetadataConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddMetadataConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.AddMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddMetadataConfiguration")
	}
}

// CallWithLogging_AddMetadataConfiguration works like Call_AddMetadataConfiguration but also logs the response body.
func CallWithLogging_AddMetadataConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.AddMetadataConfiguration) (media.AddMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddMetadataConfigurationResponse media.AddMetadataConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddMetadataConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddMetadataConfiguration")
		return reply.Body.AddMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddMetadataConfiguration")
	}
}
