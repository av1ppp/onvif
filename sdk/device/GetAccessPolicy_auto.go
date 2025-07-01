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

// Call_GetAccessPolicy forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAccessPolicyResponse.
func Call_GetAccessPolicy(ctx context.Context, dev *onvif.Device, request device.GetAccessPolicy) (device.GetAccessPolicyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAccessPolicyResponse device.GetAccessPolicyResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetAccessPolicyResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAccessPolicy")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetAccessPolicyResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAccessPolicy")
	}
	return reply.Body.GetAccessPolicyResponse, nil
}

// CallWithLogging_GetAccessPolicy works like Call_GetAccessPolicy but also logs the response body.
func CallWithLogging_GetAccessPolicy(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetAccessPolicy) (device.GetAccessPolicyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAccessPolicyResponse device.GetAccessPolicyResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetAccessPolicyResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAccessPolicy")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAccessPolicy")
	if err != nil {
		return reply.Body.GetAccessPolicyResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAccessPolicy")
	}
	return reply.Body.GetAccessPolicyResponse, nil
}
