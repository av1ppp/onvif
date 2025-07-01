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

// Call_RemoveVideoSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveVideoSourceConfigurationResponse.
func Call_RemoveVideoSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.RemoveVideoSourceConfiguration) (media.RemoveVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveVideoSourceConfigurationResponse media.RemoveVideoSourceConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RemoveVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveVideoSourceConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.RemoveVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveVideoSourceConfiguration")
	}
	return reply.Body.RemoveVideoSourceConfigurationResponse, nil
}

// CallWithLogging_RemoveVideoSourceConfiguration works like Call_RemoveVideoSourceConfiguration but also logs the response body.
func CallWithLogging_RemoveVideoSourceConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.RemoveVideoSourceConfiguration) (media.RemoveVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveVideoSourceConfigurationResponse media.RemoveVideoSourceConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RemoveVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveVideoSourceConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RemoveVideoSourceConfiguration")
	if err != nil {
		return reply.Body.RemoveVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveVideoSourceConfiguration")
	}
	return reply.Body.RemoveVideoSourceConfigurationResponse, nil
}
