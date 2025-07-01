// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// Call_AddScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a AddScopesResponse.
func Call_AddScopes(ctx context.Context, dev *onvif.Device, request device.AddScopes) (device.AddScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddScopesResponse device.AddScopesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.AddScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddScopes")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.AddScopesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddScopes")
	}
	return reply.Body.AddScopesResponse, nil
}

// CallWithLogging_AddScopes works like Call_AddScopes but also logs the response body.
func CallWithLogging_AddScopes(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.AddScopes) (device.AddScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddScopesResponse device.AddScopesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.AddScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddScopes")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddScopes")
	if err != nil {
		return reply.Body.AddScopesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddScopes")
	}
	return reply.Body.AddScopesResponse, nil
}
