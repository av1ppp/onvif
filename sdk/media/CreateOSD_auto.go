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

// Call_CreateOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateOSDResponse.
func Call_CreateOSD(ctx context.Context, dev *onvif.Device, request media.CreateOSD) (media.CreateOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateOSDResponse media.CreateOSDResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.CreateOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateOSD")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.CreateOSDResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "CreateOSD")
	}
	return reply.Body.CreateOSDResponse, nil
}

// CallWithLogging_CreateOSD works like Call_CreateOSD but also logs the response body.
func CallWithLogging_CreateOSD(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.CreateOSD) (media.CreateOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateOSDResponse media.CreateOSDResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethodWithLogging(logger, request)
	if err != nil {
		return reply.Body.CreateOSDResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "CreateOSD")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "CreateOSD")
	if err != nil {
		return reply.Body.CreateOSDResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "CreateOSD")
	}
	return reply.Body.CreateOSDResponse, nil
}
