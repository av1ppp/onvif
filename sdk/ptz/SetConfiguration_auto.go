// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_SetConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetConfigurationResponse.
func Call_SetConfiguration(ctx context.Context, dev *onvif.Device, request ptz.SetConfiguration) (ptz.SetConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetConfigurationResponse ptz.SetConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetConfiguration")
	}
	return reply.Body.SetConfigurationResponse, nil
}

// CallWithLogging_SetConfiguration works like Call_SetConfiguration but also logs the response body.
func CallWithLogging_SetConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.SetConfiguration) (ptz.SetConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetConfigurationResponse ptz.SetConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetConfiguration")
	if err != nil {
		return reply.Body.SetConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetConfiguration")
	}
	return reply.Body.SetConfigurationResponse, nil
}
