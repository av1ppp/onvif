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

// Call_GetIPAddressFilter forwards the call to dev.CallMethod() then parses the payload of the reply as a GetIPAddressFilterResponse.
func Call_GetIPAddressFilter(ctx context.Context, dev *onvif.Device, request device.GetIPAddressFilter) (device.GetIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetIPAddressFilterResponse device.GetIPAddressFilterResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetIPAddressFilterResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetIPAddressFilter")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetIPAddressFilterResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetIPAddressFilter")
	}
}

// CallWithLogging_GetIPAddressFilter works like Call_GetIPAddressFilter but also logs the response body.
func CallWithLogging_GetIPAddressFilter(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetIPAddressFilter) (device.GetIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetIPAddressFilterResponse device.GetIPAddressFilterResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetIPAddressFilterResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetIPAddressFilter")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetIPAddressFilter")
		return reply.Body.GetIPAddressFilterResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetIPAddressFilter")
	}
}
