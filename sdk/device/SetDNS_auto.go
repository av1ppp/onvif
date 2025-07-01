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

// Call_SetDNS forwards the call to dev.CallMethod() then parses the payload of the reply as a SetDNSResponse.
func Call_SetDNS(ctx context.Context, dev *onvif.Device, request device.SetDNS) (device.SetDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDNSResponse device.SetDNSResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetDNSResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetDNS")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetDNSResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetDNS")
	}
	return reply.Body.SetDNSResponse, nil
}

// CallWithLogging_SetDNS works like Call_SetDNS but also logs the response body.
func CallWithLogging_SetDNS(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetDNS) (device.SetDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDNSResponse device.SetDNSResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetDNSResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetDNS")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetDNS")
	if err != nil {
		return reply.Body.SetDNSResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetDNS")
	}
	return reply.Body.SetDNSResponse, nil
}
