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

// Call_GetDynamicDNS forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDynamicDNSResponse.
func Call_GetDynamicDNS(ctx context.Context, dev *onvif.Device, request device.GetDynamicDNS) (device.GetDynamicDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDynamicDNSResponse device.GetDynamicDNSResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDynamicDNSResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDynamicDNS")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetDynamicDNSResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDynamicDNS")
	}
}

// CallWithLogging_GetDynamicDNS works like Call_GetDynamicDNS but also logs the response body.
func CallWithLogging_GetDynamicDNS(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetDynamicDNS) (device.GetDynamicDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDynamicDNSResponse device.GetDynamicDNSResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDynamicDNSResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDynamicDNS")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDynamicDNS")
		return reply.Body.GetDynamicDNSResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDynamicDNS")
	}
}
