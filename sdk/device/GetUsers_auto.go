// Code generated : DO NOT EDIT.

package device

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/device"
	"github.com/av1ppp/onvif/errors"
)

// GetUsers forwards the call to onvif.Do then parses the payload of the reply as a GetUsersResponse.
func GetUsers(ctx context.Context, dev *onvif.Device, request *onvif.Req[device.GetUsers]) (device.GetUsersResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetUsersResponse device.GetUsersResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetUsersResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetUsers")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetUsersResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetUsers")
	}
	return reply.Body.GetUsersResponse, nil
}
