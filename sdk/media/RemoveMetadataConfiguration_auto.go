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

// Call_RemoveMetadataConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveMetadataConfigurationResponse.
func Call_RemoveMetadataConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveMetadataConfiguration) (media.RemoveMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveMetadataConfigurationResponse media.RemoveMetadataConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RemoveMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveMetadataConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.RemoveMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveMetadataConfiguration")
	}
	return reply.Body.RemoveMetadataConfigurationResponse, nil
}

// CallWithLogging_RemoveMetadataConfiguration works like Call_RemoveMetadataConfiguration but also logs the response body.
func CallWithLogging_RemoveMetadataConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.RemoveMetadataConfiguration) (media.RemoveMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveMetadataConfigurationResponse media.RemoveMetadataConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RemoveMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveMetadataConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RemoveMetadataConfiguration")
	if err != nil {
		return reply.Body.RemoveMetadataConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveMetadataConfiguration")
	}
	return reply.Body.RemoveMetadataConfigurationResponse, nil
}
