// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Stop forwards the call to onvif.Do then parses the payload of the reply as a StopResponse.
func Stop(ctx context.Context, dev *onvif.Device, request *onvif.Req[ptz.Stop]) (ptz.StopResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StopResponse ptz.StopResponse
		}
	}
	var reply Envelope

	httpReply, err := onvif.Do(dev, request)
	if err != nil {
		return reply.Body.StopResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "Stop")
	} 

	err = sdk.ReadAndParse(ctx, httpReply, &reply)
	if err != nil {
		return reply.Body.StopResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "Stop")
	}
	return reply.Body.StopResponse, nil
}
