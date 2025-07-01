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

// Call_GetConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetConfigurationsResponse.
func Call_GetConfigurations(ctx context.Context, dev *onvif.Device, request ptz.GetConfigurations) (ptz.GetConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationsResponse ptz.GetConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetConfigurations")
	}
}

// CallWithLogging_GetConfigurations works like Call_GetConfigurations but also logs the response body.
func CallWithLogging_GetConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.GetConfigurations) (ptz.GetConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationsResponse ptz.GetConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetConfigurations")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetConfigurations")
		return reply.Body.GetConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetConfigurations")
	}
}
