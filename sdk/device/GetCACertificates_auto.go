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

// Call_GetCACertificates forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCACertificatesResponse.
func Call_GetCACertificates(ctx context.Context, dev *onvif.Device, request device.GetCACertificates) (device.GetCACertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCACertificatesResponse device.GetCACertificatesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCACertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCACertificates")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.GetCACertificatesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCACertificates")
	}
}

// CallWithLogging_GetCACertificates works like Call_GetCACertificates but also logs the response body.
func CallWithLogging_GetCACertificates(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetCACertificates) (device.GetCACertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCACertificatesResponse device.GetCACertificatesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCACertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCACertificates")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCACertificates")
		return reply.Body.GetCACertificatesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCACertificates")
	}
}
