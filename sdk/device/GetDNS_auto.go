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

// Call_GetDNS forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDNSResponse.
func Call_GetDNS(ctx context.Context, dev *onvif.Device, request device.GetDNS) (device.GetDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDNSResponse device.GetDNSResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDNSResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDNS")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetDNSResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDNS")
	}
}

// CallWithLogging_GetDNS works like Call_GetDNS but also logs the response body.
func CallWithLogging_GetDNS(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetDNS) (device.GetDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDNSResponse device.GetDNSResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDNSResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetDNS")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetDNS")
		return reply.Body.GetDNSResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetDNS")
	}
}
