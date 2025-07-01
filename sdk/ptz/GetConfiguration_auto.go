// Code generated : DO NOT EDIT.

package ptz

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/ptz"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetConfigurationResponse.
func Call_GetConfiguration(ctx context.Context, dev *onvif.Device, request ptz.GetConfiguration) (ptz.GetConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationResponse ptz.GetConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetConfiguration")
		return reply.Body.GetConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetConfiguration")
	}
}
