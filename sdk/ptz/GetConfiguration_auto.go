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

// Call_GetConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetConfigurationResponse.
func Call_GetConfiguration(ctx context.Context, dev *onvif.Device, request ptz.GetConfiguration) (ptz.GetConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationResponse ptz.GetConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetConfiguration")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetConfiguration")
	}
	return reply.Body.GetConfigurationResponse, nil
}

// CallWithLogging_GetConfiguration works like Call_GetConfiguration but also logs the response body.
func CallWithLogging_GetConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.GetConfiguration) (ptz.GetConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationResponse ptz.GetConfigurationResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetConfiguration")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetConfiguration")
	if err != nil {
		return reply.Body.GetConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetConfiguration")
	}
	return reply.Body.GetConfigurationResponse, nil
}
