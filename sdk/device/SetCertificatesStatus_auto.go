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

// Call_SetCertificatesStatus forwards the call to dev.CallMethod() then parses the payload of the reply as a SetCertificatesStatusResponse.
func Call_SetCertificatesStatus(ctx context.Context, dev *onvif.Device, request device.SetCertificatesStatus) (device.SetCertificatesStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetCertificatesStatusResponse device.SetCertificatesStatusResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetCertificatesStatusResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetCertificatesStatus")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.SetCertificatesStatusResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetCertificatesStatus")
	}
}

// CallWithLogging_SetCertificatesStatus works like Call_SetCertificatesStatus but also logs the response body.
func CallWithLogging_SetCertificatesStatus(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetCertificatesStatus) (device.SetCertificatesStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetCertificatesStatusResponse device.SetCertificatesStatusResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetCertificatesStatusResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetCertificatesStatus")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetCertificatesStatus")
		return reply.Body.SetCertificatesStatusResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetCertificatesStatus")
	}
}
