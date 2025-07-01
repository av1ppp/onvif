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

// Call_GetCompatibleConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleConfigurationsResponse.
func Call_GetCompatibleConfigurations(ctx context.Context, dev *onvif.Device, request ptz.GetCompatibleConfigurations) (ptz.GetCompatibleConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleConfigurationsResponse ptz.GetCompatibleConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetCompatibleConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleConfigurations")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetCompatibleConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleConfigurations")
	}
	return reply.Body.GetCompatibleConfigurationsResponse, nil
}

// CallWithLogging_GetCompatibleConfigurations works like Call_GetCompatibleConfigurations but also logs the response body.
func CallWithLogging_GetCompatibleConfigurations(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.GetCompatibleConfigurations) (ptz.GetCompatibleConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleConfigurationsResponse ptz.GetCompatibleConfigurationsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetCompatibleConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleConfigurations")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCompatibleConfigurations")
	if err != nil {
		return reply.Body.GetCompatibleConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleConfigurations")
	}
	return reply.Body.GetCompatibleConfigurationsResponse, nil
}
