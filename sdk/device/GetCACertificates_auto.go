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

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetCACertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCACertificates")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetCACertificatesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCACertificates")
	}
	return reply.Body.GetCACertificatesResponse, nil
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

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.GetCACertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCACertificates")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetCACertificates")
	if err != nil {
		return reply.Body.GetCACertificatesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCACertificates")
	}
	return reply.Body.GetCACertificatesResponse, nil
}
