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

// Call_GetConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetConfigurationOptionsResponse.
func Call_GetConfigurationOptions(ctx context.Context, dev *onvif.Device, request ptz.GetConfigurationOptions) (ptz.GetConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationOptionsResponse ptz.GetConfigurationOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetConfigurationOptions")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetConfigurationOptions")
	}
	return reply.Body.GetConfigurationOptionsResponse, nil
}

// CallWithLogging_GetConfigurationOptions works like Call_GetConfigurationOptions but also logs the response body.
func CallWithLogging_GetConfigurationOptions(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.GetConfigurationOptions) (ptz.GetConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationOptionsResponse ptz.GetConfigurationOptionsResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetConfigurationOptions")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetConfigurationOptions")
	if err != nil {
		return reply.Body.GetConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetConfigurationOptions")
	}
	return reply.Body.GetConfigurationOptionsResponse, nil
}
