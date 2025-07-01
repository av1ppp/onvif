// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// GetSnapshotUri forwards the call to onvif.Do then parses the payload of the reply as a GetSnapshotUriResponse.
func GetSnapshotUri(ctx context.Context, dev *onvif.Device, request *onvif.Req[media.GetSnapshotUri]) (media.GetSnapshotUriResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSnapshotUriResponse media.GetSnapshotUriResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.GetSnapshotUriResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetSnapshotUri")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.GetSnapshotUriResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetSnapshotUri")
	}
	return reply.Body.GetSnapshotUriResponse, nil
}
