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

// Call_RemoveScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveScopesResponse.
func Call_RemoveScopes(ctx context.Context, dev *onvif.Device, request device.RemoveScopes) (device.RemoveScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveScopesResponse device.RemoveScopesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.RemoveScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveScopes")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.RemoveScopesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveScopes")
	}
	return reply.Body.RemoveScopesResponse, nil
}

// CallWithLogging_RemoveScopes works like Call_RemoveScopes but also logs the response body.
func CallWithLogging_RemoveScopes(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.RemoveScopes) (device.RemoveScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveScopesResponse device.RemoveScopesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.RemoveScopesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "RemoveScopes")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "RemoveScopes")
	if err != nil {
		return reply.Body.RemoveScopesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "RemoveScopes")
	}
	return reply.Body.RemoveScopesResponse, nil
}
