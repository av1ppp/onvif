// Code generated : DO NOT EDIT.

package device

import (
	"context"

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
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAccessPolicyResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAccessPolicy")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAccessPolicy")
		return reply.Body.GetAccessPolicyResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetAccessPolicy")
	}
}
