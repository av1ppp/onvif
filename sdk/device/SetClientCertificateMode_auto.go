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

// Call_SetClientCertificateMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetClientCertificateModeResponse.
func Call_SetClientCertificateMode(ctx context.Context, dev *onvif.Device, request device.SetClientCertificateMode) (device.SetClientCertificateModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetClientCertificateModeResponse device.SetClientCertificateModeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetClientCertificateModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetClientCertificateMode")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.SetClientCertificateModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetClientCertificateMode")
	}
	return reply.Body.SetClientCertificateModeResponse, nil
}

// CallWithLogging_SetClientCertificateMode works like Call_SetClientCertificateMode but also logs the response body.
func CallWithLogging_SetClientCertificateMode(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.SetClientCertificateMode) (device.SetClientCertificateModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetClientCertificateModeResponse device.SetClientCertificateModeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.SetClientCertificateModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "SetClientCertificateMode")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "SetClientCertificateMode")
	if err != nil {
		return reply.Body.SetClientCertificateModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "SetClientCertificateMode")
	}
	return reply.Body.SetClientCertificateModeResponse, nil
}
