// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// CreateUsers forwards the call to onvif.Do then parses the payload of the reply as a CreateUsersResponse.
func CreateUsers(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.CreateUsers]) (device.CreateUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateUsersResponse device.CreateUsersResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.CreateUsersResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateUsers")
	}

	logger := dev.GetLogger()
	if logger != nil {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreateUsers")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
	}
	return reply.Body.CreateUsersResponse, err
}
