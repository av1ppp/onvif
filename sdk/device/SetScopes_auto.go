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

// Call_SetScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a SetScopesResponse.
func Call_SetScopes(ctx context.Context, dev *onvif.Device, request device.SetScopes) (device.SetScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetScopesResponse device.SetScopesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetScopes")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetScopesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetScopes")
	}
	return reply.Body.SetScopesResponse, nil
}

// CallWithLogging_SetScopes works like Call_SetScopes but also logs the response body.
func CallWithLogging_SetScopes(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetScopes) (device.SetScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetScopesResponse device.SetScopesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetScopes")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetScopes")
	if err != nil {
		return reply.Body.SetScopesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetScopes")
	}
	return reply.Body.SetScopesResponse, nil
}
