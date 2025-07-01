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

// Call_GetSnapshotUri forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSnapshotUriResponse.
func Call_GetSnapshotUri(ctx context.Context, dev *onvif.Device, request media.GetSnapshotUri) (media.GetSnapshotUriResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSnapshotUriResponse media.GetSnapshotUriResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetSnapshotUriResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSnapshotUri")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetSnapshotUriResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSnapshotUri")
	}
	return reply.Body.GetSnapshotUriResponse, nil
}

// CallWithLogging_GetSnapshotUri works like Call_GetSnapshotUri but also logs the response body.
func CallWithLogging_GetSnapshotUri(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.GetSnapshotUri) (media.GetSnapshotUriResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSnapshotUriResponse media.GetSnapshotUriResponse
		}
	}
	var reply Envelope

	httpReply, err := dev.CallMethod(request)
	if err != nil {
		return reply.Body.GetSnapshotUriResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSnapshotUri")
	} 

	err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "GetSnapshotUri")
	if err != nil {
		return reply.Body.GetSnapshotUriResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSnapshotUri")
	}
	return reply.Body.GetSnapshotUriResponse, nil
}
