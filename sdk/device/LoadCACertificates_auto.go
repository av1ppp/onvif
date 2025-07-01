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

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.LoadCACertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "LoadCACertificates")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.LoadCACertificatesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "LoadCACertificates")
	}
	return reply.Body.LoadCACertificatesResponse, nil
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

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.LoadCACertificatesResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "LoadCACertificates")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "LoadCACertificates")
	if err != nil {
		return reply.Body.LoadCACertificatesResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "LoadCACertificates")
	}
	return reply.Body.LoadCACertificatesResponse, nil
}
