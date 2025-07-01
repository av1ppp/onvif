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

// Call_AddVideoSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddVideoSourceConfigurationResponse.
func Call_AddVideoSourceConfiguration(ctx context.Context, dev *onvif.Device, request media.AddVideoSourceConfiguration) (media.AddVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddVideoSourceConfigurationResponse media.AddVideoSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddVideoSourceConfiguration")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply)
		return reply.Body.AddVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddVideoSourceConfiguration")
	}
}

// CallWithLogging_AddVideoSourceConfiguration works like Call_AddVideoSourceConfiguration but also logs the response body.
func CallWithLogging_AddVideoSourceConfiguration(ctx context.Context, logger *logx.Logger, dev *onvif.Device, request media.AddVideoSourceConfiguration) (media.AddVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddVideoSourceConfigurationResponse media.AddVideoSourceConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.AddVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to call method").WithProperty(errors.PropMethod, "AddVideoSourceConfiguration")
	} else {
		err = sdk.ReadAndParseWithLogging(ctx, logger, httpReply, &reply, "AddVideoSourceConfiguration")
		return reply.Body.AddVideoSourceConfigurationResponse, errors.Common.Wrap(err, "failed to read and parse reply").WithProperty(errors.PropMethod, "AddVideoSourceConfiguration")
	}
}
