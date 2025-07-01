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

// Call_SetHostnameFromDHCP forwards the call to dev.CallMethod() then parses the payload of the reply as a SetHostnameFromDHCPResponse.
func Call_SetHostnameFromDHCP(ctx context.Context, dev *onvif.Device, request device.SetHostnameFromDHCP) (device.SetHostnameFromDHCPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHostnameFromDHCPResponse device.SetHostnameFromDHCPResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetHostnameFromDHCPResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetHostnameFromDHCP")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetHostnameFromDHCPResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetHostnameFromDHCP")
	}
	return reply.Body.SetHostnameFromDHCPResponse, nil
}

// CallWithLogging_SetHostnameFromDHCP works like Call_SetHostnameFromDHCP but also logs the response body.
func CallWithLogging_SetHostnameFromDHCP(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetHostnameFromDHCP) (device.SetHostnameFromDHCPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHostnameFromDHCPResponse device.SetHostnameFromDHCPResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetHostnameFromDHCPResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetHostnameFromDHCP")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetHostnameFromDHCP")
	if err != nil {
		return reply.Body.SetHostnameFromDHCPResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetHostnameFromDHCP")
	}
	return reply.Body.SetHostnameFromDHCPResponse, nil
}
