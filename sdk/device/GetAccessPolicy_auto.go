// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetAccessPolicy forwards the call to onvif.Do then parses the payload of the reply as a GetAccessPolicyResponse.
func GetAccessPolicy(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetAccessPolicy]) (device.GetAccessPolicyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAccessPolicyResponse device.GetAccessPolicyResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.GetAccessPolicyResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetAccessPolicy")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetAccessPolicy")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.GetAccessPolicyResponse, err
}
