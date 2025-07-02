// Code generated : DO NOT EDIT.

package sdkdevice

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// SetAccessPolicy forwards the call to onvif.Do then parses the payload of the reply as a SetAccessPolicyResponse.
func SetAccessPolicy(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.SetAccessPolicy]) (device.SetAccessPolicyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAccessPolicyResponse device.SetAccessPolicyResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.SetAccessPolicyResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetAccessPolicy")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetAccessPolicy")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.SetAccessPolicyResponse, err
}
