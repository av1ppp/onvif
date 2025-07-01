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

// Call_GetCertificatesStatus forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCertificatesStatusResponse.
func Call_GetCertificatesStatus(ctx context.Context, dev *onvif.Device, request device.GetCertificatesStatus) (device.GetCertificatesStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificatesStatusResponse device.GetCertificatesStatusResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetCertificatesStatusResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCertificatesStatus")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetCertificatesStatusResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCertificatesStatus")
	}
	return reply.Body.GetCertificatesStatusResponse, nil
}

// CallWithLogging_GetCertificatesStatus works like Call_GetCertificatesStatus but also logs the response body.
func CallWithLogging_GetCertificatesStatus(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.GetCertificatesStatus) (device.GetCertificatesStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificatesStatusResponse device.GetCertificatesStatusResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetCertificatesStatusResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCertificatesStatus")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCertificatesStatus")
	if err != nil {
		return reply.Body.GetCertificatesStatusResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCertificatesStatus")
	}
	return reply.Body.GetCertificatesStatusResponse, nil
}
