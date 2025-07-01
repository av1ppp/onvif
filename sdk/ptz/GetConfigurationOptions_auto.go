// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetConfigurationOptionsResponse.
func Call_GetConfigurationOptions(ctx context.Context, dev *onvif.Device, request ptz.GetConfigurationOptions) (ptz.GetConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationOptionsResponse ptz.GetConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetConfigurationOptions")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetConfigurationOptions")
		return reply.Body.GetConfigurationOptionsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetConfigurationOptions")
	}
}
