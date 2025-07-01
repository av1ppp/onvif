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

// Call_SetVideoSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoSourceConfigurationResponse.
func Call_SetVideoSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.SetVideoSourceConfiguration) (media.SetVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoSourceConfigurationResponse media.SetVideoSourceConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetVideoSourceConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetVideoSourceConfiguration")
	}
	return reply.Body.SetVideoSourceConfigurationResponse, nil
}

// CallWithLogging_SetVideoSourceConfiguration works like Call_SetVideoSourceConfiguration but also logs the response body.
func CallWithLogging_SetVideoSourceConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.SetVideoSourceConfiguration) (media.SetVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoSourceConfigurationResponse media.SetVideoSourceConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetVideoSourceConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetVideoSourceConfiguration")
	if err != nil {
		return reply.Body.SetVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetVideoSourceConfiguration")
	}
	return reply.Body.SetVideoSourceConfigurationResponse, nil
}
