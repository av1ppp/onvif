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

// Call_GetClientCertificateMode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetClientCertificateModeResponse.
func Call_GetClientCertificateMode(ctx context.Context, dev *onvif.Device, request device.GetClientCertificateMode) (device.GetClientCertificateModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetClientCertificateModeResponse device.GetClientCertificateModeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetClientCertificateModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetClientCertificateMode")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetClientCertificateModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetClientCertificateMode")
	}
	return reply.Body.GetClientCertificateModeResponse, nil
}

// CallWithLogging_GetClientCertificateMode works like Call_GetClientCertificateMode but also logs the response body.
func CallWithLogging_GetClientCertificateMode(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetClientCertificateMode) (device.GetClientCertificateModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetClientCertificateModeResponse device.GetClientCertificateModeResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetClientCertificateModeResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetClientCertificateMode")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetClientCertificateMode")
	if err != nil {
		return reply.Body.GetClientCertificateModeResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetClientCertificateMode")
	}
	return reply.Body.GetClientCertificateModeResponse, nil
}
