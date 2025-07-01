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

// Call_SetDynamicDNS forwards the call to dev.CallMethod() then parses the payload of the reply as a SetDynamicDNSResponse.
func Call_SetDynamicDNS(ctx context.Context, dev *onvif.Device, request device.SetDynamicDNS) (device.SetDynamicDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDynamicDNSResponse device.SetDynamicDNSResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetDynamicDNSResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetDynamicDNS")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetDynamicDNSResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetDynamicDNS")
	}
	return reply.Body.SetDynamicDNSResponse, nil
}

// CallWithLogging_SetDynamicDNS works like Call_SetDynamicDNS but also logs the response body.
func CallWithLogging_SetDynamicDNS(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetDynamicDNS) (device.SetDynamicDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDynamicDNSResponse device.SetDynamicDNSResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.SetDynamicDNSResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetDynamicDNS")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetDynamicDNS")
	if err != nil {
		return reply.Body.SetDynamicDNSResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetDynamicDNS")
	}
	return reply.Body.SetDynamicDNSResponse, nil
}
