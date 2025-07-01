// Code generated : DO NOT EDIT.

package media

import (
	"context"

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
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSnapshotUriResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSnapshotUri")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetSnapshotUri")
		return reply.Body.GetSnapshotUriResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSnapshotUri")
	}
}
