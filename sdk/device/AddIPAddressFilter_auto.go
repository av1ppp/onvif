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

// Call_AddIPAddressFilter forwards the call to dev.CallMethod() then parses the payload of the reply as a AddIPAddressFilterResponse.
func Call_AddIPAddressFilter(ctx context.Context, dev *onvif.Device, request device.AddIPAddressFilter) (device.AddIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddIPAddressFilterResponse device.AddIPAddressFilterResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.AddIPAddressFilterResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddIPAddressFilter")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.AddIPAddressFilterResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddIPAddressFilter")
	}
	return reply.Body.AddIPAddressFilterResponse, nil
}

// CallWithLogging_AddIPAddressFilter works like Call_AddIPAddressFilter but also logs the response body.
func CallWithLogging_AddIPAddressFilter(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.AddIPAddressFilter) (device.AddIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddIPAddressFilterResponse device.AddIPAddressFilterResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.AddIPAddressFilterResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddIPAddressFilter")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddIPAddressFilter")
	if err != nil {
		return reply.Body.AddIPAddressFilterResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddIPAddressFilter")
	}
	return reply.Body.AddIPAddressFilterResponse, nil
}
