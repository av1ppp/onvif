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

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetConfigurations")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetConfigurations")
	}
	return reply.Body.GetConfigurationsResponse, nil
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

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetConfigurations")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetConfigurations")
	if err != nil {
		return reply.Body.GetConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetConfigurations")
	}
	return reply.Body.GetConfigurationsResponse, nil
}
