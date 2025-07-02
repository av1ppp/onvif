// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// DeleteUsers forwards the call to onvif.Do then parses the payload of the reply as a DeleteUsersResponse.
func DeleteUsers(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.DeleteUsers]) (device.DeleteUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteUsersResponse device.DeleteUsersResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(ctx, dev, request)
	if err != nil {
		return reply.Body.DeleteUsersResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteUsers")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "DeleteUsers")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.DeleteUsersResponse, err
}
