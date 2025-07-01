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

// Call_SetMetadataConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetMetadataConfigurationResponse.
func Call_SetMetadataConfiguration(ctx context.Context, dev *onvif.Device, request media.SetMetadataConfiguration) (media.SetMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetMetadataConfigurationResponse media.SetMetadataConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetMetadataConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetMetadataConfiguration")
	}
	return reply.Body.SetMetadataConfigurationResponse, nil
}

// CallWithLogging_SetMetadataConfiguration works like Call_SetMetadataConfiguration but also logs the response body.
func CallWithLogging_SetMetadataConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.SetMetadataConfiguration) (media.SetMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetMetadataConfigurationResponse media.SetMetadataConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetMetadataConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetMetadataConfiguration")
	if err != nil {
		return reply.Body.SetMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetMetadataConfiguration")
	}
	return reply.Body.SetMetadataConfigurationResponse, nil
}
