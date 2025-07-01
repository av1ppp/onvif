// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/logx"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_DeleteOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteOSDResponse.
func Call_DeleteOSD(ctx context.Context, dev *onvif.Device, request media.DeleteOSD) (media.DeleteOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteOSDResponse media.DeleteOSDResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.DeleteOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteOSD")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.DeleteOSDResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteOSD")
	}
	return reply.Body.DeleteOSDResponse, nil
}

// CallWithLogging_DeleteOSD works like Call_DeleteOSD but also logs the response body.
func CallWithLogging_DeleteOSD(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.DeleteOSD) (media.DeleteOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteOSDResponse media.DeleteOSDResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.DeleteOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "DeleteOSD")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "DeleteOSD")
	if err != nil {
		return reply.Body.DeleteOSDResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "DeleteOSD")
	}
	return reply.Body.DeleteOSDResponse, nil
}
