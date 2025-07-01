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

// Call_LoadCACertificates forwards the call to dev.CallMethod() then parses the payload of the reply as a LoadCACertificatesResponse.
func Call_LoadCACertificates(ctx context.Context, dev *onvif.Device, request device.LoadCACertificates) (device.LoadCACertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			LoadCACertificatesResponse device.LoadCACertificatesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.LoadCACertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "LoadCACertificates")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.LoadCACertificatesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "LoadCACertificates")
	}
}

// CallWithLogging_LoadCACertificates works like Call_LoadCACertificates but also logs the response body.
func CallWithLogging_LoadCACertificates(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.LoadCACertificates) (device.LoadCACertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			LoadCACertificatesResponse device.LoadCACertificatesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.LoadCACertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "LoadCACertificates")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "LoadCACertificates")
		return reply.Body.LoadCACertificatesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "LoadCACertificates")
	}
}
