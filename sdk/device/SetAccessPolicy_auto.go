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

// Call_SetAccessPolicy forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAccessPolicyResponse.
func Call_SetAccessPolicy(ctx context.Context, dev *onvif.Device, request device.SetAccessPolicy) (device.SetAccessPolicyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAccessPolicyResponse device.SetAccessPolicyResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetAccessPolicyResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAccessPolicy")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetAccessPolicyResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetAccessPolicy")
	}
	return reply.Body.SetAccessPolicyResponse, nil
}

// CallWithLogging_SetAccessPolicy works like Call_SetAccessPolicy but also logs the response body.
func CallWithLogging_SetAccessPolicy(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetAccessPolicy) (device.SetAccessPolicyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAccessPolicyResponse device.SetAccessPolicyResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetAccessPolicyResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAccessPolicy")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetAccessPolicy")
	if err != nil {
		return reply.Body.SetAccessPolicyResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetAccessPolicy")
	}
	return reply.Body.SetAccessPolicyResponse, nil
}
