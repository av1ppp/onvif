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

// Call_AddPTZConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddPTZConfigurationResponse.
func Call_AddPTZConfiguration(ctx context.Context, dev *onvif.Device, request media.AddPTZConfiguration) (media.AddPTZConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddPTZConfigurationResponse media.AddPTZConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddPTZConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddPTZConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.AddPTZConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddPTZConfiguration")
	}
}

// CallWithLogging_AddPTZConfiguration works like Call_AddPTZConfiguration but also logs the response body.
func CallWithLogging_AddPTZConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.AddPTZConfiguration) (media.AddPTZConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddPTZConfigurationResponse media.AddPTZConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddPTZConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddPTZConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddPTZConfiguration")
		return reply.Body.AddPTZConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddPTZConfiguration")
	}
}
