// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetStatus forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStatusResponse.
func Call_GetStatus(ctx context.Context, dev *onvif.Device, request ptz.GetStatus) (ptz.GetStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStatusResponse ptz.GetStatusResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetStatusResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetStatus")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetStatusResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetStatus")
	}
	return reply.Body.GetStatusResponse, nil
}

// CallWithLogging_GetStatus works like Call_GetStatus but also logs the response body.
func CallWithLogging_GetStatus(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request ptz.GetStatus) (ptz.GetStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStatusResponse ptz.GetStatusResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetStatusResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetStatus")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetStatus")
	if err != nil {
		return reply.Body.GetStatusResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetStatus")
	}
	return reply.Body.GetStatusResponse, nil
}
