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

// Call_GetVideoSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoSourceConfigurationResponse.
func Call_GetVideoSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.GetVideoSourceConfiguration) (media.GetVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceConfigurationResponse media.GetVideoSourceConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoSourceConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoSourceConfiguration")
	}
	return reply.Body.GetVideoSourceConfigurationResponse, nil
}

// CallWithLogging_GetVideoSourceConfiguration works like Call_GetVideoSourceConfiguration but also logs the response body.
func CallWithLogging_GetVideoSourceConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetVideoSourceConfiguration) (media.GetVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceConfigurationResponse media.GetVideoSourceConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetVideoSourceConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetVideoSourceConfiguration")
	if err != nil {
		return reply.Body.GetVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetVideoSourceConfiguration")
	}
	return reply.Body.GetVideoSourceConfigurationResponse, nil
}
