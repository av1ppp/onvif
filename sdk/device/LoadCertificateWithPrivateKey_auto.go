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

// Call_LoadCertificateWithPrivateKey forwards the call to dev.CallMethod() then parses the payload of the reply as a LoadCertificateWithPrivateKeyResponse.
func Call_LoadCertificateWithPrivateKey(ctx context.Context, dev *onvif.Device, request device.LoadCertificateWithPrivateKey) (device.LoadCertificateWithPrivateKeyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			LoadCertificateWithPrivateKeyResponse device.LoadCertificateWithPrivateKeyResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.LoadCertificateWithPrivateKeyResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "LoadCertificateWithPrivateKey")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.LoadCertificateWithPrivateKeyResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "LoadCertificateWithPrivateKey")
	}
	return reply.Body.LoadCertificateWithPrivateKeyResponse, nil
}

// CallWithLogging_LoadCertificateWithPrivateKey works like Call_LoadCertificateWithPrivateKey but also logs the response body.
func CallWithLogging_LoadCertificateWithPrivateKey(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request device.LoadCertificateWithPrivateKey) (device.LoadCertificateWithPrivateKeyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			LoadCertificateWithPrivateKeyResponse device.LoadCertificateWithPrivateKeyResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.LoadCertificateWithPrivateKeyResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "LoadCertificateWithPrivateKey")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "LoadCertificateWithPrivateKey")
	if err != nil {
		return reply.Body.LoadCertificateWithPrivateKeyResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "LoadCertificateWithPrivateKey")
	}
	return reply.Body.LoadCertificateWithPrivateKeyResponse, nil
}
