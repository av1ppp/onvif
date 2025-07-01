// Code generated : DO NOT EDIT.

package media

import (
	"context"

	"github.com/av1ppp/onvif"
	"github.com/av1ppp/onvif/sdk"
	"github.com/av1ppp/onvif/media"
	"github.com/av1ppp/onvif/errors"
)

// Call_GetCompatibleVideoSourceConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleVideoSourceConfigurationsResponse.
func Call_GetCompatibleVideoSourceConfigurations(ctx context.Context, dev *onvif.Device, request media.GetCompatibleVideoSourceConfigurations) (media.GetCompatibleVideoSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoSourceConfigurationsResponse media.GetCompatibleVideoSourceConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleVideoSourceConfigurationsResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "GetCompatibleVideoSourceConfigurations")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCompatibleVideoSourceConfigurations")
		return reply.Body.GetCompatibleVideoSourceConfigurationsResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "GetCompatibleVideoSourceConfigurations")
	}
}
