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

// Call_GetDPAddresses forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDPAddressesResponse.
func Call_GetDPAddresses(ctx context.Context, dev *onvif.Device, request device.GetDPAddresses) (device.GetDPAddressesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDPAddressesResponse device.GetDPAddressesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetDPAddressesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDPAddresses")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetDPAddressesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDPAddresses")
	}
	return reply.Body.GetDPAddressesResponse, nil
}

// CallWithLogging_GetDPAddresses works like Call_GetDPAddresses but also logs the response body.
func CallWithLogging_GetDPAddresses(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetDPAddresses) (device.GetDPAddressesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDPAddressesResponse device.GetDPAddressesResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetDPAddressesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDPAddresses")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDPAddresses")
	if err != nil {
		return reply.Body.GetDPAddressesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDPAddresses")
	}
	return reply.Body.GetDPAddressesResponse, nil
}
